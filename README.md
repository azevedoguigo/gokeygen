# gokeygen

`gokeygen` is a simple and lightweight Go library for generating random keys.  
It allows you to create numeric, alphabetic, or alphanumeric keys with support for uppercase and lowercase formats.

Ideal for:

- Tokens
- Temporary IDs
- Verification codes
- Access keys
- Non-cryptographic seeds
- Testing and mocks

> ⚠️ This library uses `math/rand` and **is NOT suitable for cryptographic use**.

---

## Installation

```bash
go get github.com/azevedoguigo/gokeygen
```

## Usage Examples

```go
package main

import (
	"fmt"

	"github.com/azevedoguigo/gokeygen"
)

func main() {
	key := gokeygen.GenerateNumericKey(6)
    fmt.Println(key) //Result: 483920

	key := gokeygen.GenerateLettersKey(8)
    fmt.Println(key) //Result: aZbYxKqW

	key := gokeygen.GenerateLettersKeyUppercase(8)
    fmt.Println(key) //Result: QWERTYUI

	key := gokeygen.GenerateLettersKeyLowercase(8)
    fmt.Println(key) //Result: asdfghjk

    key := gokeygen.GenerateAlphanumericKey(10)
    fmt.Println(key) //Result: a8K2pL9xQ1

    key := gokeygen.GenerateAlphanumericKeyUppercase(10)
    fmt.Println(key) //Result: A8K2PL9XQ1

	key := gokeygen.GenerateAlphanumericKeyLowecase(10)
    fmt.Println(key) //Result: a8k2pl9xq1
}
```
