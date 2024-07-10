// Package network provides the implementation for managing nodes in the blockchain network.
package network

import "blockchain/blockchain"

// Node represents a node in the blockchain network.
type Node struct {
	Address string
}

// Network represents the blockchain network containing multiple nodes.
type Network struct {
	Nodes []Node
}

// AddNode adds a new node with the given address to the network.
func (n *Network) AddNode(address string) {
	n.Nodes = append(n.Nodes, Node{Address: address})
}

// GetNodes returns all the nodes in the network.
func (n *Network) GetNodes() []Node {
	return n.Nodes
}

// SyncBlockchain synchronizes the blockchain across all nodes in the network.
// This function needs to be implemented.
func (n *Network) SyncBlockchain(blocks []blockchain.Block) {
	// Implement blockchain synchronization between nodes
}
