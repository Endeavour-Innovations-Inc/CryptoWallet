// main.go
package main

import (
    "fmt"
    "cryptowallet/crypto"
    "cryptowallet/wallet"
)

func main() {
    privateKey, publicKey := wallet.GenerateKeyPair()
    fmt.Println("Private Key:", privateKey)
    fmt.Println("Public Key:", publicKey)

    // For demonstration, let's encrypt and save the private key
    encryptedPrivateKey := crypto.Encrypt([]byte(privateKey), "your_secure_passphrase")
    err := crypto.SaveToFile("wallet.dat", encryptedPrivateKey)
    if err != nil {
        fmt.Println("Error saving wallet:", err)
        return
    }
    fmt.Println("Encrypted wallet saved successfully.")
}
