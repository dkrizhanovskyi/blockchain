package api

import (
	"net/http"
	"blockchain/internal/blockchain"
	"blockchain/internal/utils"
)

// RegisterRoutes sets up the API routes and their corresponding handlers.
// Parameters:
// - blockchain: The blockchain instance to be used by the handlers.
// - logger: The logger instance for logging API activities.
// Returns:
// - An http.ServeMux that maps the routes to their handlers.
func RegisterRoutes(blockchain *blockchain.Blockchain, logger *utils.Logger) *http.ServeMux {
	// Create a new ServeMux to register the routes.
	mux := http.NewServeMux()

	// Initialize the handlers with the provided blockchain and logger.
	handlers := NewHandlers(blockchain, logger)

	// Register the route for adding a new block.
	mux.HandleFunc("/addblock", handlers.AddBlockHandler)

	// Register the route for getting the entire blockchain.
	mux.HandleFunc("/getblockchain", handlers.GetBlockchainHandler)

	// Register the route for getting a specific block by index.
	mux.HandleFunc("/block", handlers.GetBlockByIndexHandler)

	// Register the route for getting the last block.
	mux.HandleFunc("/lastblock", handlers.GetLastBlockHandler)

	// Register the route for validating the blockchain.
	mux.HandleFunc("/validate", handlers.ValidateBlockchainHandler)

	// Return the configured ServeMux.
	return mux
}
