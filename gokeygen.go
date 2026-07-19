package gokeygen

import (
	"math/rand"
)

const charsetLettersLower = "abcdefghijklmnopqrstuvwxyz"
const charsetLettersUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetLetters = charsetLettersLower + charsetLettersUpper
const charsetNumeric = "0123456789"
const charsetAlphanumeric = charsetLetters + charsetNumeric
const charsetAlphanumericLower = charsetLettersLower + charsetNumeric
const charsetAlphanumericUpper = charsetLettersUpper + charsetNumeric

func generateKey(charset string, length int) string {
	if length <= 0 {
		return ""
	}

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
	return generateKey(charsetLettersUpper, length)
}

func GenerateLettersKeyLowercase(length int) string {
	return generateKey(charsetLettersLower, length)
}

func GenerateAlphanumericKey(length int) string {
	return generateKey(charsetAlphanumeric, length)
}

func GenerateAlphanumericKeyUppercase(length int) string {
	return generateKey(charsetAlphanumericUpper, length)
}

func GenerateAlphanumericKeyLowercase(length int) string {
	return generateKey(charsetAlphanumericLower, length)
}
