package blockchain

import (
	"testing"
)

// TestAddBlock tests the AddBlock function to ensure that blocks are added correctly to the blockchain.
func TestAddBlock(t *testing.T) {
	bc := GetBlockchain("SHA-256") // Initialize a new blockchain.

	// Add a new block to the blockchain.
	bc.AddBlock("Test Block 1")

	// Check that the blockchain has two blocks (the genesis block + the new block).
	if len(bc.Blocks) != 2 {
		t.Errorf("Expected 2 blocks, but got %d", len(bc.Blocks))
	}

	// Verify the data of the new block.
	if bc.Blocks[1].Data != "Test Block 1" {
		t.Errorf("Expected block data 'Test Block 1', but got '%s'", bc.Blocks[1].Data)
	}
}

// TestAddBlockWithRust tests the AddBlockWithRust function to ensure that blocks are added correctly using Rust's hashing implementation.
func TestAddBlockWithRust(t *testing.T) {
	bc := GetBlockchain("SHA-256") // Initialize a new blockchain.

	// Add a new block using Rust's hashing implementation.
	bc.AddBlockWithRust("Test Block 2")

	// Check that the blockchain has two blocks (the genesis block + the new block).
	if len(bc.Blocks) != 2 {
		t.Errorf("Expected 2 blocks, but got %d", len(bc.Blocks))
	}

	// Verify the data of the new block.
	if bc.Blocks[1].Data != "Test Block 2" {
		t.Errorf("Expected block data 'Test Block 2', but got '%s'", bc.Blocks[1].Data)
	}
}

// TestIsChainValid tests the IsChainValid function to ensure that the blockchain integrity is correctly validated.
func TestIsChainValid(t *testing.T) {
	bc := GetBlockchain("SHA-256") // Initialize a new blockchain.

	// Add a couple of blocks.
	bc.AddBlock("Test Block 1")
	bc.AddBlock("Test Block 2")

	// Verify that the blockchain is valid.
	if !bc.IsChainValid() {
		t.Error("Expected blockchain to be valid, but it is not.")
	}

	// Manually tamper with the blockchain to simulate an invalid state.
	bc.Blocks[1].Data = "Tampered Data"

	// Verify that the blockchain is now detected as invalid.
	if bc.IsChainValid() {
		t.Error("Expected blockchain to be invalid, but it is still considered valid.")
	}
}
