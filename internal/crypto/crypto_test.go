package crypto

import (
	"encoding/hex"
	"testing"
)

// TestHashSHA256Go tests the SHA-256 hashing function implemented in Go.
func TestHashSHA256Go(t *testing.T) {
	input := []byte("test data") // Sample input data.
	expectedHash := "916f0027a575074ce72a331777c3478d6513f786a591bd892da1a577bf2335f9" // Expected SHA-256 hash.

	// Compute the hash using the Go-based function.
	output := HashSHA256Go(input)

	// Convert the output hash to a hexadecimal string.
	outputHex := hex.EncodeToString(output)

	// Compare the computed hash with the expected value.
	if outputHex != expectedHash {
		t.Errorf("Expected SHA-256 hash %s, but got %s", expectedHash, outputHex)
	}
}

// TestHashSHA512Go tests the SHA-512 hashing function implemented in Go.
func TestHashSHA512Go(t *testing.T) {
	input := []byte("test data") // Sample input data.
	expectedHash := "0e1e21ecf105ec853d24d728867ad70613c21663a4693074b2a3619c1bd39d66b588c33723bb466c72424e80e3ca63c249078ab347bab9428500e7ee43059d0d" // Expected SHA-512 hash.

	// Compute the hash using the Go-based function.
	output := HashSHA512Go(input)

	// Convert the output hash to a hexadecimal string.
	outputHex := hex.EncodeToString(output)

	// Compare the computed hash with the expected value.
	if outputHex != expectedHash {
		t.Errorf("Expected SHA-512 hash %s, but got %s", expectedHash, outputHex)
	}
}
