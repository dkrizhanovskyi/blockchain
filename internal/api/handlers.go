package api

import (
	"blockchain/internal/blockchain"
	"encoding/json"
	"net/http"
	"strconv"
	"blockchain/internal/utils"
)

// Handlers struct holds the blockchain instance and logger to be used by the API handlers.
type Handlers struct {
	Blockchain *blockchain.Blockchain
	Logger     *utils.Logger
}

// NewHandlers creates a new Handlers instance.
// Parameters:
// - blockchain: The blockchain instance to interact with.
// - logger: The logger instance for logging API activities.
// Returns:
// - A new Handlers instance.
func NewHandlers(blockchain *blockchain.Blockchain, logger *utils.Logger) *Handlers {
	return &Handlers{
		Blockchain: blockchain,
		Logger:     logger,
	}
}

// AddBlockHandler handles the API request to add a new block to the blockchain.
// This is a POST request handler.
// It expects a JSON body with a "data" field.
func (h *Handlers) AddBlockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		h.Logger.Error("Failed to decode request body:", err)
		return
	}

	h.Blockchain.AddBlockWithRust(req.Data)
	h.Logger.Info("New block added with data:", req.Data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Block added successfully"})
}

// GetBlockchainHandler handles the API request to get the entire blockchain.
// This is a GET request handler.
func (h *Handlers) GetBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.Blockchain.Blocks); err != nil {
		http.Error(w, "Failed to encode blockchain data", http.StatusInternalServerError)
		h.Logger.Error("Failed to encode blockchain data:", err)
		return
	}
	h.Logger.Info("Blockchain data retrieved")
}

// GetBlockByIndexHandler handles the API request to get a specific block by its index.
// This is a GET request handler.
// It expects an "index" query parameter.
func (h *Handlers) GetBlockByIndexHandler(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")
	if indexStr == "" {
		http.Error(w, "Index is required", http.StatusBadRequest)
		h.Logger.Warn("Index query parameter missing")
		return
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		h.Logger.Error("Invalid index format:", err)
		return
	}

	if index < 0 || index >= len(h.Blockchain.Blocks) {
		http.Error(w, "Index out of range", http.StatusBadRequest)
		h.Logger.Warn("Index out of range:", index)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.Blockchain.Blocks[index]); err != nil {
		http.Error(w, "Failed to encode block data", http.StatusInternalServerError)
		h.Logger.Error("Failed to encode block data:", err)
		return
	}
	h.Logger.Info("Block data retrieved for index:", index)
}

// GetLastBlockHandler handles the API request to get the last block in the blockchain.
// This is a GET request handler.
func (h *Handlers) GetLastBlockHandler(w http.ResponseWriter, r *http.Request) {
	lastBlock := h.Blockchain.Blocks[len(h.Blockchain.Blocks)-1]

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lastBlock); err != nil {
		http.Error(w, "Failed to encode block data", http.StatusInternalServerError)
		h.Logger.Error("Failed to encode last block data:", err)
		return
	}
	h.Logger.Info("Last block data retrieved")
}

// ValidateBlockchainHandler handles the API request to validate the blockchain.
// This is a GET request handler.
func (h *Handlers) ValidateBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	valid := h.Blockchain.IsChainValid()

	response := map[string]string{
		"message": "Blockchain is valid",
	}
	if !valid {
		response["message"] = "Blockchain is invalid"
		w.WriteHeader(http.StatusConflict)
		h.Logger.Warn("Blockchain validation failed")
	} else {
		h.Logger.Info("Blockchain validation successful")
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode validation result", http.StatusInternalServerError)
		h.Logger.Error("Failed to encode validation result:", err)
		return
	}
}
