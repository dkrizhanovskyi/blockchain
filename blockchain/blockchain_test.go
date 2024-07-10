// blockchain/blockchain_test.go
package blockchain

import (
	"blockchain/crypto"
	"fmt"
	"testing"
	"time"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()
	if len(bc.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(bc.Blocks))
	}
	if bc.Blocks[0].Index != 0 {
		t.Fatalf("Expected genesis block index to be 0, got %d", bc.Blocks[0].Index)
	}
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockchain()

	// Generate key pair for the validator
	privKey, pubKey, err := crypto.GenerateKeyPair()
	if err != nil {
		t.Fatalf("Error generating key pair: %v", err)
	}
	pubKeyStr, err := crypto.ExportPublicKey(pubKey)
	if err != nil {
		t.Fatalf("Error exporting public key: %v", err)
	}
	bc.Validators[pubKeyStr] = 100

	tx := Transaction{Sender: pubKeyStr, Recipient: "bob", Amount: 10}
	signature, err := crypto.Sign(privKey, fmt.Sprintf("%s%s%d", tx.Sender, tx.Recipient, tx.Amount))
	if err != nil {
		t.Fatalf("Error signing transaction: %v", err)
	}
	tx.Signature = signature

	bc.AddBlock([]Transaction{tx}, pubKeyStr, 100)

	if len(bc.Blocks) != 2 {
		t.Fatalf("Expected 2 blocks, got %d", len(bc.Blocks))
	}
	if bc.Blocks[1].Index != 1 {
		t.Fatalf("Expected new block index to be 1, got %d", bc.Blocks[1].Index)
	}
}

func TestSelectValidator(t *testing.T) {
	bc := NewBlockchain()
	bc.Validators["validator1"] = 100
	bc.Validators["validator2"] = 200

	validator := bc.SelectValidator()
	if validator != "validator1" && validator != "validator2" {
		t.Fatalf("Expected validator1 or validator2, got %s", validator)
	}
}

func TestCalculateHash(t *testing.T) {
	block := Block{
		Index:     1,
		Timestamp: time.Now().String(),
		Transactions: []Transaction{
			{Sender: "alice", Recipient: "bob", Amount: 10},
		},
		PrevHash:  "prevhash",
		Validator: "validator",
		Stake:     100,
	}
	hash := CalculateHash(block)
	if hash == "" {
		t.Fatalf("Expected non-empty hash, got empty hash")
	}
}
