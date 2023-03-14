package main

import (
	"crypto/rand"
	"fmt"
)

func generateCryptoRandom(chars string, length int32) string {
	bytes := make([]byte, length)
	rand.Read(bytes)

	for index, element := range bytes {
		randomize := element % byte(len(chars))
		bytes[index] = chars[randomize]
	}

	return string(bytes)
}

func main() {
	fmt.Println(generateCryptoRandom("abcdefghijklmnopqrstuvwxyz0123456789", 30))
}
