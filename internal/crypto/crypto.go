package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
)

// HashSHA256Go computes the SHA-256 hash of the input data using Go's built-in crypto package.
// Parameters:
// - input: The input data to be hashed.
// Returns:
// - The resulting 32-byte SHA-256 hash as a byte slice.
func HashSHA256Go(input []byte) []byte {
	hash := sha256.Sum256(input) // Compute the SHA-256 hash of the input data.
	return hash[:]               // Return the hash as a byte slice.
}

// HashSHA512Go computes the SHA-512 hash of the input data using Go's built-in crypto package.
// Parameters:
// - input: The input data to be hashed.
// Returns:
// - The resulting 64-byte SHA-512 hash as a byte slice.
func HashSHA512Go(input []byte) []byte {
	hash := sha512.Sum512(input) // Compute the SHA-512 hash of the input data.
	return hash[:]               // Return the hash as a byte slice.
}
