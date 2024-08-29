package crypto

/*
#cgo LDFLAGS: -L../../rust/rust_crypto/target/release -lrust_crypto

// Function declarations for linking with the Rust shared library.
extern void hash_sha256(const unsigned char* input, unsigned long len, unsigned char* output);
extern void hash_sha512(const unsigned char* input, unsigned long len, unsigned char* output);
*/
import "C"
import (
	"unsafe"
)

// HashSHA256 computes the SHA-256 hash of the input data using the Rust implementation.
// Parameters:
// - input: The input data to be hashed.
// Returns:
// - The resulting 32-byte SHA-256 hash.
func HashSHA256(input []byte) []byte {
	output := make([]byte, 32) // Prepare a buffer to hold the 32-byte SHA-256 hash.

	// Call the Rust function via C bindings.
	// Convert the Go byte slice to a C pointer and pass it to the Rust function.
	C.hash_sha256((*C.uchar)(unsafe.Pointer(&input[0])), C.ulong(len(input)), (*C.uchar)(unsafe.Pointer(&output[0])))

	return output // Return the computed hash.
}

// HashSHA512 computes the SHA-512 hash of the input data using the Rust implementation.
// Parameters:
// - input: The input data to be hashed.
// Returns:
// - The resulting 64-byte SHA-512 hash.
func HashSHA512(input []byte) []byte {
	output := make([]byte, 64) // Prepare a buffer to hold the 64-byte SHA-512 hash.

	// Call the Rust function via C bindings.
	// Convert the Go byte slice to a C pointer and pass it to the Rust function.
	C.hash_sha512((*C.uchar)(unsafe.Pointer(&input[0])), C.ulong(len(input)), (*C.uchar)(unsafe.Pointer(&output[0])))

	return output // Return the computed hash.
}
