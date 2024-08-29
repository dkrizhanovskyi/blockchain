package p2p

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/network"
	"blockchain/internal/utils"
	"log"
	"net"
)

// Node represents a single node in the P2P network.
type Node struct {
	Blockchain *blockchain.Blockchain // The blockchain instance managed by the node.
	Network    *network.Network       // The network instance to handle peer connections.
	Logger     *utils.Logger          // Logger for logging node activities.
	Address    string                 // The address this node is listening on.
}

// NewNode creates and initializes a new Node instance.
// Parameters:
// - address: The address this node will listen on.
// - blockchain: The blockchain instance managed by the node.
// - logger: The logger instance for logging node activities.
// Returns:
// - A new Node instance.
func NewNode(address string, blockchain *blockchain.Blockchain, logger *utils.Logger) *Node {
	return &Node{
		Blockchain: blockchain,
		Network:    network.NewNetwork(),
		Logger:     logger,
		Address:    address,
	}
}

// Start starts the node's server to listen for incoming connections from peers.
func (n *Node) Start() {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", n.Address, err)
	}
	defer listener.Close()

	n.Logger.Info("Node is listening on", n.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			n.Logger.Error("Failed to accept connection:", err)
			continue
		}

		peer := &network.Peer{
			Address: conn.RemoteAddr().String(),
			Conn:    conn,
		}

		n.Network.Peers[peer.Address] = peer
		go n.Network.HandleConnection(peer)
	}
}

// ConnectToPeer connects this node to another peer in the network.
// Parameters:
// - address: The address of the peer to connect to.
func (n *Node) ConnectToPeer(address string) {
	err := n.Network.ConnectToPeer(address)
	if err != nil {
		n.Logger.Error("Failed to connect to peer:", err)
	} else {
		n.Logger.Info("Connected to peer:", address)
	}
}

// BroadcastBlock broadcasts a new block to all connected peers.
// Parameters:
// - block: The block to be broadcasted.
func (n *Node) BroadcastBlock(block *blockchain.Block) {
	message := "BLOCK " + block.Data
	n.Network.Broadcast(message)
	n.Logger.Info("Broadcasted block with data:", block.Data)
}

// HandleMessage processes an incoming message and takes appropriate action.
// Parameters:
// - message: The message received from a peer.
func (n *Node) HandleMessage(message string) {
	n.Logger.Info("Received message:", message)

	// This is where different types of messages can be handled.
	// For example, if a message contains a new block, add it to the blockchain.
	if message[:6] == "BLOCK " {
		blockData := message[6:]
		n.Blockchain.AddBlockWithRust(blockData)
		n.Logger.Info("New block added with data:", blockData)
	} else {
		n.Logger.Warn("Unknown message type received:", message)
	}
}
