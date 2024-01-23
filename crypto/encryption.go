// encryption.go
package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "os"

    "golang.org/x/crypto/sha3"
)

func Encrypt(data []byte, passphrase string) []byte {
    key := sha3.Sum256([]byte(passphrase))
    block, err := aes.NewCipher(key[:])
    if err != nil {
        panic(err) // or handle the error as per your application's error policy
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }

    return gcm.Seal(nonce, nonce, data, nil)
}

func SaveToFile(filename string, data []byte) error {
    return os.WriteFile(filename, data, 0644)
}
