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
	"runtime/pprof"
	"sync"
	"syscall"
	"time"
)

var cache sync.Map

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Генерация ключей
	privKey, pubKey, err := crypto.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	// Экспорт публичного ключа в строку
	pubKeyStr, err := crypto.ExportPublicKey(pubKey)
	if err != nil {
		fmt.Println("Error exporting public key:", err)
		return
	}

	// Установка системного публичного ключа
	blockchain.SystemPublicKey = pubKeyStr

	// Создание блокчейна с генезис-блоком и начальной ставкой валидатора
	bc := blockchain.NewBlockchain()
	bc.Validators[pubKeyStr] = 100 // Начальная ставка валидатора

	// Инициализация API для взаимодействия с блокчейном и сетью узлов
	go api.InitAPIWithCache(bc, &cache)

	// Обслуживание статических файлов
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Настройка TLS для защиты соединений
	server := &http.Server{
		Addr: ":8081",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	// Запуск автоматического создания блоков каждые 30 секунд
	go startAutomaticBlockCreation(bc, privKey, pubKeyStr)

	// Обработка завершения работы сервера
	go handleShutdown(server)

	// Запуск сервера с использованием TLS
	fmt.Println("Starting server on :8081")
	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func startAutomaticBlockCreation(bc *blockchain.Blockchain, privKey *ecdsa.PrivateKey, pubKeyStr string) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		createAutomaticBlock(bc, privKey, pubKeyStr)
		cache.Store("blockchain", bc.Blocks)
	}
}

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
	cache.Store("blockchain", bc.Blocks) // Кеширование состояния блокчейна
}

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
