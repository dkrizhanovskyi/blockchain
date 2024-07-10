package network

import (
	"blockchain/blockchain"
	"sync"
)

// Node represents a node in the network.
type Node struct {
	Address string
}

// Network represents the network of nodes.
type Network struct {
	Nodes []Node
	mu    sync.Mutex
}

// AddNode adds a new node to the network.
func (n *Network) AddNode(address string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Nodes = append(n.Nodes, Node{Address: address})
}

// GetNodes returns all nodes in the network.
func (n *Network) GetNodes() []Node {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.Nodes
}

// SyncBlockchain synchronizes the blockchain with all nodes.
func (n *Network) SyncBlockchain(blocks []blockchain.Block) {
	// Implement blockchain synchronization logic here
}
