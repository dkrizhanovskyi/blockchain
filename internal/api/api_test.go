package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"blockchain/internal/blockchain"
	"blockchain/internal/utils"
)

// TestAddBlockHandler tests the AddBlockHandler to ensure it correctly adds a block to the blockchain.
func TestAddBlockHandler(t *testing.T) {
	logger := utils.NewLogger("Test: ", 0)
	bc := blockchain.GetBlockchain("SHA-256")
	handlers := NewHandlers(bc, logger)

	// Create a sample request body with block data.
	body := map[string]string{"data": "Test Block"}
	bodyBytes, _ := json.Marshal(body)

	// Create a new HTTP POST request.
	req, err := http.NewRequest("POST", "/addblock", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddBlockHandler)

	// Serve the request.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message":"Block added successfully"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Ensure that the blockchain now contains two blocks (the genesis block + the new block).
	if len(bc.Blocks) != 2 {
		t.Errorf("Expected 2 blocks in the blockchain, but got %d", len(bc.Blocks))
	}
}

// TestGetBlockchainHandler tests the GetBlockchainHandler to ensure it returns the correct blockchain data.
func TestGetBlockchainHandler(t *testing.T) {
	logger := utils.NewLogger("Test: ", 0)
	bc := blockchain.GetBlockchain("SHA-256")
	handlers := NewHandlers(bc, logger)

	// Add a block to the blockchain for testing.
	bc.AddBlock("Test Block")

	// Create a new HTTP GET request.
	req, err := http.NewRequest("GET", "/getblockchain", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBlockchainHandler)

	// Serve the request.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body contains the expected blockchain data.
	var blocks []*blockchain.Block
	if err := json.Unmarshal(rr.Body.Bytes(), &blocks); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if len(blocks) != 2 {
		t.Errorf("Expected 2 blocks in the blockchain, but got %d", len(blocks))
	}

	if blocks[1].Data != "Test Block" {
		t.Errorf("Expected block data 'Test Block', but got '%s'", blocks[1].Data)
	}
}

// TestGetBlockByIndexHandler tests the GetBlockByIndexHandler to ensure it returns the correct block by index.
func TestGetBlockByIndexHandler(t *testing.T) {
	logger := utils.NewLogger("Test: ", 0)
	bc := blockchain.GetBlockchain("SHA-256")
	handlers := NewHandlers(bc, logger)

	// Add a block to the blockchain for testing.
	bc.AddBlock("Test Block")

	// Create a new HTTP GET request with the index parameter.
	req, err := http.NewRequest("GET", "/block?index=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBlockByIndexHandler)

	// Serve the request.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body contains the expected block data.
	var block blockchain.Block
	if err := json.Unmarshal(rr.Body.Bytes(), &block); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if block.Data != "Test Block" {
		t.Errorf("Expected block data 'Test Block', but got '%s'", block.Data)
	}
}

// TestGetLastBlockHandler tests the GetLastBlockHandler to ensure it returns the last block in the blockchain.
func TestGetLastBlockHandler(t *testing.T) {
	logger := utils.NewLogger("Test: ", 0)
	bc := blockchain.GetBlockchain("SHA-256")
	handlers := NewHandlers(bc, logger)

	// Add a block to the blockchain for testing.
	bc.AddBlock("Test Block")

	// Create a new HTTP GET request.
	req, err := http.NewRequest("GET", "/lastblock", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetLastBlockHandler)

	// Serve the request.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body contains the expected last block data.
	var block blockchain.Block
	if err := json.Unmarshal(rr.Body.Bytes(), &block); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if block.Data != "Test Block" {
		t.Errorf("Expected last block data 'Test Block', but got '%s'", block.Data)
	}
}

// TestValidateBlockchainHandler tests the ValidateBlockchainHandler to ensure it correctly validates the blockchain.
func TestValidateBlockchainHandler(t *testing.T) {
	logger := utils.NewLogger("Test: ", 0)
	bc := blockchain.GetBlockchain("SHA-256")
	handlers := NewHandlers(bc, logger)

	// Add a couple of blocks to the blockchain for testing.
	bc.AddBlock("Test Block 1")
	bc.AddBlock("Test Block 2")

	// Create a new HTTP GET request.
	req, err := http.NewRequest("GET", "/validate", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ValidateBlockchainHandler)

	// Serve the request.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body contains the expected validation message.
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if response["message"] != "Blockchain is valid" {
		t.Errorf("Expected validation message 'Blockchain is valid', but got '%s'", response["message"])
	}

	// Tamper with the blockchain to make it invalid.
	bc.Blocks[1].Data = "Tampered Data"

	// Serve the request again to validate the tampered blockchain.
	rr = httptest.NewRecorder()  // Reset the response recorder
	handler.ServeHTTP(rr, req)

	// Unmarshal the response again to check the updated validation status.
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response body after tampering: %v", err)
	}

	// Check that the blockchain is now invalid.
	if response["message"] != "Blockchain is invalid" {
		t.Errorf("Expected validation message 'Blockchain is invalid', but got '%s'", response["message"])
	}
}
