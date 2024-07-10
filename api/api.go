package api

import (
	"blockchain/blockchain"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// User represents a user in the system.
type User struct {
	Username string
	Password string
}

// InitAPIWithCache initializes the API with caching.
func InitAPIWithCache(bc *blockchain.Blockchain, cacheInstance *sync.Map) {
	fmt.Println("Initializing API...")
	blockchainInstance = bc
	cache = cacheInstance

	http.HandleFunc("/blockchain", handleGetBlockchainWithCache)
	http.HandleFunc("/write", handleWriteBlock)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/send", handleSendTransaction)
	http.HandleFunc("/mine", handleMineNewCoins)

	fmt.Println("Starting API server on :8082")
	err := http.ListenAndServeTLS(":8082", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("Error starting API server:", err)
	}
}

func handleGetBlockchainWithCache(w http.ResponseWriter, r *http.Request) {
	if cachedBlocks, ok := cache.Load("blockchain"); ok {
		bytes, err := json.MarshalIndent(cachedBlocks, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		return
	}

	bytes, err := json.MarshalIndent(blockchainInstance.Blocks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	cache.Store("blockchain", blockchainInstance.Blocks)
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var transactions []blockchain.Transaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transactions); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	defer mutex.Unlock()
	transactionBuffer = append(transactionBuffer, transactions...)
	if len(transactionBuffer) >= maxBufferSize {
		validator := blockchainInstance.SelectValidator()
		if validator == "" {
			fmt.Println("No validator selected")
			return
		}
		blockchainInstance.AddBlock(transactionBuffer, validator, blockchainInstance.Validators[validator])
		transactionBuffer = []blockchain.Transaction{}
		cache.Store("blockchain", blockchainInstance.Blocks)
	}

	respondWithJSON(w, http.StatusCreated, blockchainInstance.Blocks)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	passwordHash := sha256.Sum256([]byte(req.Password))
	user := User{
		Username:     req.Username,
		PasswordHash: hex.EncodeToString(passwordHash[:]),
	}

	users.Lock()
	users.m[req.Username] = user
	users.Unlock()

	respondWithJSON(w, http.StatusCreated, "User registered successfully")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	users.RLock()
	user, exists := users.m[req.Username]
	users.RUnlock()

	hashedPassword := sha256.Sum256([]byte(req.Password))
	if !exists || user.PasswordHash != hex.EncodeToString(hashedPassword[:]) {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	sessionToken := make([]byte, 32)
	rand.Read(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   hex.EncodeToString(sessionToken),
		Expires: time.Now().Add(24 * time.Hour),
	})

	respondWithJSON(w, http.StatusOK, "Login successful")
}

func handleSendTransaction(w http.ResponseWriter, r *http.Request) {
	type SendTransactionRequest struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
		Signature string `json:"signature"`
	}

	var req SendTransactionRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	transaction := blockchain.Transaction{
		Sender:    req.Sender,
		Recipient: req.Recipient,
		Amount:    req.Amount,
		Signature: req.Signature,
	}

	mutex.Lock()
	defer mutex.Unlock()
	transactionBuffer = append(transactionBuffer, transaction)
	if len(transactionBuffer) >= maxBufferSize {
		validator := blockchainInstance.SelectValidator()
		if validator == "" {
			fmt.Println("No validator selected")
			return
		}
		blockchainInstance.AddBlock(transactionBuffer, validator, blockchainInstance.Validators[validator])
		transactionBuffer = []blockchain.Transaction{}
		cache.Store("blockchain", blockchainInstance.Blocks)
	}

	respondWithJSON(w, http.StatusCreated, "Transaction sent successfully")
}

func handleMineNewCoins(w http.ResponseWriter, r *http.Request) {
	type MineRequest struct {
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	}

	var mineReq MineRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mineReq); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	mutex.Lock()
	blockchainInstance.MineNewCoins(mineReq.Recipient, mineReq.Amount)
	cache.Store("blockchain", blockchainInstance.Blocks)
	mutex.Unlock()

	respondWithJSON(w, http.StatusCreated, blockchainInstance.Blocks)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
