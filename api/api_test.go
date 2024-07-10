package api

import (
	"blockchain/blockchain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// assertEqualBlocks checks if two blocks are equal, ignoring dynamic fields like Timestamp and Hash.
func assertEqualBlocks(t *testing.T, got, want blockchain.Block) {
	if got.Index != want.Index {
		t.Errorf("got Index %v, want %v", got.Index, want.Index)
	}
	if got.PrevHash != want.PrevHash {
		t.Errorf("got PrevHash %v, want %v", got.PrevHash, want.PrevHash)
	}
	if got.Validator != want.Validator {
		t.Errorf("got Validator %v, want %v", got.Validator, want.Validator)
	}
	if got.Stake != want.Stake {
		t.Errorf("got Stake %v, want %v", got.Stake, want.Stake)
	}
	if len(got.Transactions) != len(want.Transactions) {
		t.Errorf("got %v transactions, want %v transactions", len(got.Transactions), len(want.Transactions))
	}
	// Further checks can be added as needed
}

func TestHandleGetBlockchain(t *testing.T) {
	// Initialize the blockchain and cache
	bc := blockchain.NewBlockchain()
	cache = &sync.Map{}
	cache.Store("blockchain", bc.Blocks)
	blockchainInstance = bc

	req, err := http.NewRequest("GET", "/blockchain", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetBlockchainWithCache)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotBlocks []blockchain.Block
	err = json.Unmarshal(rr.Body.Bytes(), &gotBlocks)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if len(gotBlocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(gotBlocks))
	}

	wantBlock := blockchain.Block{
		Index:        0,
		Timestamp:    "2024-07-09",
		Transactions: nil,
		PrevHash:     "",
		Hash:         "",
		Validator:    "",
		Stake:        0,
	}

	assertEqualBlocks(t, gotBlocks[0], wantBlock)
}
