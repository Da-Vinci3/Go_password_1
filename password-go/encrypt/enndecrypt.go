package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// PKCS7 padding
func addPadding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

// Remove PKCS7 padding
func removePadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding - data is empty")
	}

	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("invalid padding size")
	}

	for i := length - padding; i < length; i++ {
		if data[i] != byte(padding) {
			return nil, errors.New("invalid padding values")
		}
	}

	return data[:length-padding], nil
}

func EncryptAES(key []byte, plaintext string) (string, error) {
	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad the plaintext
	plainBytes := []byte(plaintext)
	plainBytes = addPadding(plainBytes, aes.BlockSize)

	// Create IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// Create CBC mode encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the data
	ciphertext := make([]byte, len(plainBytes))
	mode.CryptBlocks(ciphertext, plainBytes)

	// Prepend IV to ciphertext
	final := append(iv, ciphertext...)

	// Convert to hex string
	return hex.EncodeToString(final), nil
}

func DecryptAES(key []byte, ciphertext string) (string, error) {
	// Decode hex string
	cipherBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Check if data is too short
	if len(cipherBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Extract IV
	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	// Create CBC mode decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the data
	mode.CryptBlocks(cipherBytes, cipherBytes)

	// Remove padding
	plainBytes, err := removePadding(cipherBytes)
	if err != nil {
		return "", err
	}

	return string(plainBytes), nil
}
