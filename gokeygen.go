package gokeygen

import (
	"math/rand"
)

const charsetNumeric = "0123456789"

func GenerateNumeric(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charsetNumeric[rand.Intn(len(charsetNumeric))]
	}

	return string(result)
}
