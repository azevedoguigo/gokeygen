package gokeygen

import (
	"math/rand"
	"strings"
)

const charsetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"
const charsetAlphanumeric = charsetLetters + charsetNumeric

func generateKey(charset string, length int) string {
	bytesArr := make([]byte, length)

	for i := range bytesArr {
		bytesArr[i] = charset[rand.Intn(len(charset))]
	}

	return string(bytesArr)
}

func GenerateNumericKey(length int) string {
	return generateKey(charsetNumeric, length)
}

func GenerateLettersKey(length int) string {
	return generateKey(charsetLetters, length)
}

func GenerateLettersKeyUppercase(length int) string {
	return strings.ToUpper(generateKey(charsetLetters, length))
}

func GenerateLettersKeyLowercase(length int) string {
	return strings.ToLower(generateKey(charsetLetters, length))
}

func GenerateAlphanumericKey(length int) string {
	return generateKey(charsetAlphanumeric, length)
}

func GenerateAlphanumericKeyUppercase(length int) string {
	return strings.ToUpper(generateKey(charsetAlphanumeric, length))
}

func GenerateAlphanumericKeyLowecase(length int) string {
	return strings.ToLower(generateKey(charsetAlphanumeric, length))
}
