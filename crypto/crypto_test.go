package crypto

import (
	"testing"
)

// TestGenerateKeyPair проверяет генерацию ключевой пары.
func TestGenerateKeyPair(t *testing.T) {
	privKey, pubKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}
	if privKey == nil || pubKey == nil {
		t.Fatalf("Expected non-nil keys")
	}
}

// TestExportImportPublicKey проверяет экспорт и импорт публичного ключа.
func TestExportImportPublicKey(t *testing.T) {
	_, pubKey, _ := GenerateKeyPair()
	pubKeyStr, err := ExportPublicKey(pubKey)
	if err != nil {
		t.Fatalf("Failed to export public key: %v", err)
	}

	importedPubKey, err := ImportPublicKey(pubKeyStr)
	if err != nil {
		t.Fatalf("Failed to import public key: %v", err)
	}

	if !pubKey.Equal(importedPubKey) {
		t.Fatalf("Expected imported public key to equal original")
	}
}

// TestSignAndVerify проверяет создание и верификацию подписи.
func TestSignAndVerify(t *testing.T) {
	privKey, pubKey, _ := GenerateKeyPair()
	data := "test data"
	signature, err := Sign(privKey, data)
	if err != nil {
		t.Fatalf("Failed to sign data: %v", err)
	}

	valid, err := VerifySignature(pubKey, data, signature)
	if err != nil {
		t.Fatalf("Failed to verify signature: %v", err)
	}
	if !valid {
		t.Fatalf("Expected signature to be valid")
	}
}
