package password

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

func GeneratePass(length int) string {
	passwordSpace := make([]byte, 95)
	var bytepassword []byte

	for i := 0; i < 26; i++ {
		passwordSpace[i] = byte('a' + i)
	}

	for i := 0; i < 26; i++ {
		passwordSpace[26+i] = byte('A' + i)
	}

	for i := 0; i < 10; i++ {
		passwordSpace[52+i] = byte('0' + i)
	}

	specialChars := []byte("!@#$%^&*()-_+={}[]|\\:;\"'<>,./?`~")
	for i := 0; i < 32; i++ {
		passwordSpace[62+i] = specialChars[i]
	}

	for i := 0; i < length; i++ {
		choice := randomChoice(passwordSpace)
		bytepassword = append(bytepassword, choice)
	}

	return string(bytepassword)
}
