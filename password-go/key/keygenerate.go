package key

import (
	"crypto/rand"
	"math/big"
)

func randomChoice(list []byte) byte {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		panic(err)
	}
	return list[n.Int64()]
}

func GenerateKey(length int) string {
	keySpace := make([]byte, 95)
	var bytekey []byte

	for i := 0; i < 26; i++ {
		keySpace[i] = byte('a' + i)
	}

	for i := 0; i < 26; i++ {
		keySpace[26+i] = byte('A' + i)
	}

	for i := 0; i < 10; i++ {
		keySpace[52+i] = byte('0' + i)
	}

	specialChars := []byte("!@#$%^&*()-_+={}[]|\\:;\"'<>,./?`~")
	for i := 0; i < 32; i++ {
		keySpace[62+i] = specialChars[i]
	}

	for i := 0; i < length; i++ {
		choice := randomChoice(keySpace)
		bytekey = append(bytekey, choice)
	}

	return string(bytekey)
}
