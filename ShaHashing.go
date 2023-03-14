package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func shaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func main() {
	fmt.Println(shaHashing("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(shaHashing("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
