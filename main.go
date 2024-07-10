// Package main implements the main entry point for the blockchain server.
package main

import (
	"blockchain/api"
	"blockchain/blockchain"
	"blockchain/crypto"
	"crypto/ecdsa"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var cache sync.Map

// main is the entry point for the blockchain server.
func main() {
	// Generate keys
	privKey, pubKey, err := crypto.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	// Export public key to string
	pubKeyStr, err := crypto.ExportPublicKey(pubKey)
	if err != nil {
		fmt.Println("Error exporting public key:", err)
		return
	}

	// Set system public key
	blockchain.SystemPublicKey = pubKeyStr

	// Create a new blockchain with a genesis block and initial validator stake
	bc := blockchain.NewBlockchain()
	bc.Validators[pubKeyStr] = 100 // Initial validator stake

	// Initialize API to interact with the blockchain and node network
	go api.InitAPIWithCache(bc, &cache)

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Set up TLS to secure connections
	server := &http.Server{
		Addr: ":8081",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	// Start automatic block creation every 30 seconds
	go startAutomaticBlockCreation(bc, privKey, pubKeyStr)

	// Handle server shutdown
	go handleShutdown(server)

	// Start the server using TLS
	fmt.Println("Starting server on :8081")
	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// startAutomaticBlockCreation starts the process of creating blocks automatically every 30 seconds.
func startAutomaticBlockCreation(bc *blockchain.Blockchain, privKey *ecdsa.PrivateKey, pubKeyStr string) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		createAutomaticBlock(bc, privKey, pubKeyStr)
		cache.Store("blockchain", bc.Blocks)
	}
}

// createAutomaticBlock creates a new block with automatic transactions.
func createAutomaticBlock(bc *blockchain.Blockchain, privKey *ecdsa.PrivateKey, pubKeyStr string) {
	validator := bc.SelectValidator()
	if validator == "" {
		fmt.Println("No validator selected")
		return
	}

	transactions := []blockchain.Transaction{
		{Sender: pubKeyStr, Recipient: "automated", Amount: 0},
	}

	for i, tx := range transactions {
		signature, err := crypto.Sign(privKey, fmt.Sprintf("%s%s%d", tx.Sender, tx.Recipient, tx.Amount))
		if err != nil {
			fmt.Printf("Error signing transaction: %v\n", err)
			continue
		}
		transactions[i].Signature = signature
	}

	bc.AddBlock(transactions, validator, bc.Validators[validator])
	cache.Store("blockchain", bc.Blocks) // Cache the blockchain state
}

// handleShutdown handles server shutdown gracefully.
func handleShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Printf("Received signal: %s. Shutting down...\n", sig)

	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

	os.Exit(0)
}
