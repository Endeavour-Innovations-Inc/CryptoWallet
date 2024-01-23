// wallet.go
package wallet

import (
    "crypto/ecdsa"
    "encoding/hex"
    "log"

    "github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeyPair() (string, string) {
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
    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

    return hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes)
}

// go run wallet.go
// generates public and private key. 