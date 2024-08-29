package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt" // Import the fmt package for formatting
	"log"
	"time"

	"blockchain/internal/crypto"
)

// Block represents a single block in the blockchain.
type Block struct {
	Timestamp    int64  // The timestamp when the block was created.
	Data         string // The actual data stored in the block (e.g., transactions).
	PreviousHash string // The hash of the previous block in the chain.
	Hash         string // The hash of the current block.
}

// Blockchain represents the entire chain of blocks.
type Blockchain struct {
	Blocks []*Block // A slice holding all blocks in the blockchain.
}

// NewBlock creates a new block and computes its hash.
func NewBlock(data string, previousHash string) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
		Hash:         "",
	}
	block.Hash = block.calculateHash() // Calculate the hash for the new block.
	return block
}

// calculateHash generates the hash for a block based on its contents.
func (b *Block) calculateHash() string {
	// Use fmt.Sprintf to convert the int64 Timestamp to a string of digits.
	record := fmt.Sprintf("%d", b.Timestamp) + b.Data + b.PreviousHash
	hash := sha256.Sum256([]byte(record)) // Use SHA-256 to generate the hash.
	return hex.EncodeToString(hash[:])
}

// AddBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1] // Get the last block in the chain.
	newBlock := NewBlock(data, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock) // Append the new block to the chain.
	log.Printf("New block added: %s", newBlock.Hash)
}

// AddBlockWithRust creates a new block using Rust's hashing functions and adds it to the blockchain.
func (bc *Blockchain) AddBlockWithRust(data string) {
    previousBlock := bc.Blocks[len(bc.Blocks)-1] // Get the last block in the chain.

    // Concatenate previous hash and data to match Rust's expected input format.
    combinedData := previousBlock.Hash + data
    log.Printf("Data to be hashed by Rust: %s", combinedData)

    // Calculate the hash using Rust's SHA-256 implementation.
    hashBytes := crypto.HashSHA256([]byte(combinedData))
    hash := hex.EncodeToString(hashBytes)

    newBlock := &Block{
        Timestamp:    time.Now().Unix(),
        Data:         data,
        PreviousHash: previousBlock.Hash,
        Hash:         hash,
    }

    bc.Blocks = append(bc.Blocks, newBlock) // Append the new block to the chain.
    log.Printf("New block added using Rust: %s", newBlock.Hash)
}

// IsChainValid verifies the integrity of the blockchain.
func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Recalculate the hash and check if it matches.
		if currentBlock.Hash != currentBlock.calculateHash() {
			log.Printf("Invalid block hash at block %d", i)
			return false
		}

		// Check if the current block's PreviousHash matches the previous block's hash.
		if currentBlock.PreviousHash != previousBlock.Hash {
			log.Printf("Invalid previous hash at block %d", i)
			return false
		}
	}
	return true
}

// GetBlockchain initializes a new blockchain with a genesis block.
func GetBlockchain(hashMethod string) *Blockchain {
	genesisBlock := NewBlock("Genesis Block", "0xGENESIS")
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}
