package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	fmt.Println(md5Hash)
	//[195 252 211 215 97 146 228 0 125 251 73 108 202 103 225 59]
	// MD5 is 128 bit, which is 16 bytes
	// convert to hex based string representation
	return hex.EncodeToString(md5Hash[:])
}

func main() {
	fmt.Println(mdHashing("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(mdHashing("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
