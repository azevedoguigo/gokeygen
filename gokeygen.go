package gokeygen

import (
	"math/rand"
)

const charsetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"

func GenerateNumeric(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charsetNumeric[rand.Intn(len(charsetNumeric))]
	}

	return string(result)
}

func GenerateLetters(length int) (string, error) {
	b := make([]byte, length)

	for i := range b {
		b[i] = charsetLetters[rand.Intn(len(charsetLetters))]
	}

	return string(b), nil
}
