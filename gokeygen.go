package gokeygen

import (
	"math/rand"
	"strings"
)

const charsetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"
const charsetAlphanumeric = charsetLetters + charsetNumeric

func GenerateNumericKey(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charsetNumeric[rand.Intn(len(charsetNumeric))]
	}

	return string(result)
}

func GenerateLettersKey(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charsetLetters[rand.Intn(len(charsetLetters))]
	}

	return string(b)
}

func GenerateLettersKeyUppercase(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charsetLetters[rand.Intn(len(charsetLetters))]
	}

	return strings.ToUpper(string(b))
}

func GenerateAlphanumericKey(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charsetAlphanumeric[rand.Intn(len(charsetAlphanumeric))]
	}

	return string(b)
}
