package main

import (
    "crypto/ecdsa"
    "encoding/hex"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/crypto"
)

func generateKeyPair() {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatalf("Failed to generate private key: %v", err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Error casting public key to ECDSA")
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)
    fmt.Println("Private Key:", hex.EncodeToString(privateKeyBytes))

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
    fmt.Println("Public Key:", hex.EncodeToString(publicKeyBytes))
}

func main() {
    generateKeyPair()
}

// go run wallet.go
// generates public and private key. 