package network

import (
	"bufio"
	"log"
	"net"
)

// Peer represents a connection to another node in the network.
type Peer struct {
	Address string   // The address of the peer node.
	Conn    net.Conn // The network connection to the peer.
}

// Network manages all peer-to-peer connections for a node.
type Network struct {
	Peers map[string]*Peer // A map of connected peers, keyed by their address.
}

// NewNetwork creates and initializes a new Network instance.
func NewNetwork() *Network {
	return &Network{
		Peers: make(map[string]*Peer),
	}
}

// ConnectToPeer connects to another node in the network and adds it to the list of peers.
// Parameters:
// - address: The address of the peer to connect to.
func (n *Network) ConnectToPeer(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	peer := &Peer{
		Address: address,
		Conn:    conn,
	}

	n.Peers[address] = peer

	go n.HandleConnection(peer) // Start handling the connection in a new goroutine.

	log.Printf("Connected to peer: %s\n", address)
	return nil
}

// Broadcast sends a message to all connected peers.
// Parameters:
// - message: The message to broadcast to all peers.
func (n *Network) Broadcast(message string) {
	for _, peer := range n.Peers {
		_, err := peer.Conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Printf("Failed to send message to peer %s: %v", peer.Address, err)
		}
	}
}

// HandleConnection handles incoming messages from a peer.
// Parameters:
// - peer: The peer connection to handle.
func (n *Network) HandleConnection(peer *Peer) {
	defer peer.Conn.Close() // Ensure the connection is closed when done.
	reader := bufio.NewReader(peer.Conn)

	for {
		// Read messages from the peer, one line at a time.
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Connection closed with peer: %s", peer.Address)
			delete(n.Peers, peer.Address) // Remove the peer from the list if the connection is closed.
			return
		}

		log.Printf("Received message from peer %s: %s", peer.Address, message)
		n.HandleMessage(message)
	}
}

// HandleMessage processes an incoming message from a peer.
// Parameters:
// - message: The message received from a peer.
func (n *Network) HandleMessage(message string) {
	// This is where you would handle different types of messages.
	// For now, we'll just log the received message.
	log.Printf("Handling message: %s", message)
}
