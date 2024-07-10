package blockchain

import (
	"blockchain/crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// Blockchain represents the blockchain with a list of blocks and validators.
type Blockchain struct {
	Blocks     []Block
	Validators map[string]int
}

// Block represents a single block in the blockchain.
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
	Validator    string
	Stake        int
}

// Transaction represents a single transaction in a block.
type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
	Signature string
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:     "",
		Hash:         "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return &Blockchain{
		Blocks:     []Block{genesisBlock},
		Validators: make(map[string]int),
	}
}

// AddBlock adds a new block to the blockchain if the transactions are valid.
func (bc *Blockchain) AddBlock(transactions []Transaction, validator string, stake int) {
	if !bc.validateTransactions(transactions) {
		fmt.Println("Invalid transactions")
		return
	}
	mu.Lock()
	defer mu.Unlock()
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(prevBlock, transactions, validator, stake)
	if isBlockValid(newBlock, prevBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		bc.Validators[validator] += stake
	}
}

// SelectValidator selects a validator based on their stake.
func (bc *Blockchain) SelectValidator() string {
	totalStake := 0
	for _, stake := range bc.Validators {
		totalStake += stake
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPoint := rng.Intn(totalStake)

	currentStake := 0
	for validator, stake := range bc.Validators {
		currentStake += stake
		if currentStake > randomPoint {
			return validator
		}
	}

	return ""
}

// validateTransactions validates a list of transactions.
func (bc *Blockchain) validateTransactions(transactions []Transaction) bool {
	for _, tx := range transactions {
		if !bc.isValidTransaction(tx) {
			fmt.Printf("Invalid transaction: %+v\n", tx)
			return false
		}
	}
	return true
}

// isValidTransaction checks if a transaction is valid.
func (bc *Blockchain) isValidTransaction(tx Transaction) bool {
	if tx.Sender == "system" {
		return true
	}

	pubKey, err := crypto.ImportPublicKey(tx.Sender)
	if err != nil {
		fmt.Printf("Error importing public key: %v\n", err)
		return false
	}
	isValid, err := crypto.VerifySignature(pubKey, fmt.Sprintf("%s%s%d", tx.Sender, tx.Recipient, tx.Amount), tx.Signature)
	if err != nil || !isValid {
		fmt.Printf("Error verifying signature: %v, isValid: %t\n", err, isValid)
		return false
	}
	return true
}

// isBlockValid checks if a block is valid by comparing it to the previous block.
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// CalculateHash calculates the hash of a block.
func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%d%s%s", block.Index, block.Timestamp, block.Stake, block.Validator, block.PrevHash)
	for _, tx := range block.Transactions {
		record += tx.Sender + tx.Recipient + fmt.Sprintf("%d", tx.Amount) + tx.Signature
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CreateBlock creates a new block.
func CreateBlock(prevBlock Block, transactions []Transaction, validator string, stake int) Block {
	block := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Validator:    validator,
		Stake:        stake,
	}
	block.Hash = CalculateHash(block)
	return block
}

// MineNewCoins mines new coins by creating a new block with a transaction from the system to the recipient.
func (bc *Blockchain) MineNewCoins(recipient string, amount int) {
	transactions := []Transaction{
		{Sender: "system", Recipient: recipient, Amount: amount},
	}
	validator := bc.SelectValidator()
	bc.AddBlock(transactions, validator, bc.Validators[validator])
}
