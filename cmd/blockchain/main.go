package main

import (
	"blockchain/internal/api"
	"blockchain/internal/blockchain"
	"blockchain/internal/p2p"
	"blockchain/internal/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a logger for the application.
	logger := utils.NewLogger("BlockchainApp: ", log.LstdFlags)

	// Initialize the blockchain.
	bc := blockchain.GetBlockchain("SHA-256")

	// Create the P2P node.
	nodeAddress := getEnv("NODE_ADDRESS", "localhost:3001")
	node := p2p.NewNode(nodeAddress, bc, logger)

	// Start the P2P node in a separate goroutine.
	go node.Start()

	// Connect to any initial peers (optional, could be configured via environment or flags).
	initialPeer := getEnv("INITIAL_PEER", "")
	if initialPeer != "" {
		node.ConnectToPeer(initialPeer)
	}

	// Register API routes.
	mux := api.RegisterRoutes(bc, logger)

	// Start the HTTP server for the API.
	apiAddress := getEnv("API_ADDRESS", "localhost:8080")
	logger.Info("Starting API server on", apiAddress)
	if err := http.ListenAndServe(apiAddress, mux); err != nil {
		logger.Error("Failed to start API server:", err)
	}
}

// getEnv retrieves an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
