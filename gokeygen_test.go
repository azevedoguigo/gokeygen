package gokeygen

import "testing"

func TestGenerateKeyLength(t *testing.T) {
	generators := map[string]func(int) string{
		"Numeric":               GenerateNumericKey,
		"Letters":               GenerateLettersKey,
		"LettersUppercase":      GenerateLettersKeyUppercase,
		"LettersLowercase":      GenerateLettersKeyLowercase,
		"Alphanumeric":          GenerateAlphanumericKey,
		"AlphanumericUppercase": GenerateAlphanumericKeyUppercase,
		"AlphanumericLowercase": GenerateAlphanumericKeyLowercase,
	}

	for name, generate := range generators {
		t.Run(name, func(t *testing.T) {
			for _, length := range []int{1, 6, 32} {
				key := generate(length)
				if len(key) != length {
					t.Errorf("expected length %d, got %d (key: %q)", length, len(key), key)
				}
			}
		})
	}
}

func TestGenerateKeyNonPositiveLength(t *testing.T) {
	generators := []func(int) string{
		GenerateNumericKey,
		GenerateLettersKey,
		GenerateLettersKeyUppercase,
		GenerateLettersKeyLowercase,
		GenerateAlphanumericKey,
		GenerateAlphanumericKeyUppercase,
		GenerateAlphanumericKeyLowercase,
	}

	for _, generate := range generators {
		for _, length := range []int{0, -1, -10} {
			if key := generate(length); key != "" {
				t.Errorf("expected empty string for length %d, got %q", length, key)
			}
		}
	}
}

func TestGenerateKeyCharset(t *testing.T) {
	cases := []struct {
		name    string
		gen     func(int) string
		charset string
	}{
		{"Numeric", GenerateNumericKey, charsetNumeric},
		{"Letters", GenerateLettersKey, charsetLetters},
		{"LettersUppercase", GenerateLettersKeyUppercase, charsetLettersUpper},
		{"LettersLowercase", GenerateLettersKeyLowercase, charsetLettersLower},
		{"Alphanumeric", GenerateAlphanumericKey, charsetAlphanumeric},
		{"AlphanumericUppercase", GenerateAlphanumericKeyUppercase, charsetAlphanumericUpper},
		{"AlphanumericLowercase", GenerateAlphanumericKeyLowercase, charsetAlphanumericLower},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			key := c.gen(64)
			for _, ch := range key {
				if !containsRune(c.charset, ch) {
					t.Errorf("character %q not in expected charset %q", ch, c.charset)
				}
			}
		})
	}
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
