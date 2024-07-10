// Package crypto provides cryptographic functions for key generation, signing, and signature verification.
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
)

// GenerateKeyPair generates a new ECDSA private and public key pair.
func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privKey, &privKey.PublicKey, nil
}

// ExportPublicKey exports the given ECDSA public key to a PEM encoded string.
func ExportPublicKey(pubKey *ecdsa.PublicKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pubKeyPEM), nil
}

// ImportPublicKey imports a PEM encoded ECDSA public key string and returns the public key.
func ImportPublicKey(pubKeyStr string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*ecdsa.PublicKey), nil
}

// Sign generates a signature for the given data using the provided ECDSA private key.
func Sign(privKey *ecdsa.PrivateKey, data string) (string, error) {
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), nil
}

// VerifySignature verifies the signature of the given data using the provided ECDSA public key.
func VerifySignature(pubKey *ecdsa.PublicKey, data, signature string) (bool, error) {
	hash := sha256.Sum256([]byte(data))
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])
	return ecdsa.Verify(pubKey, hash[:], r, s), nil
}
