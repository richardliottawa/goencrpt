package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Load the private key from the file
	privateKeyFile, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		fmt.Println("Error loading private key file:", err)
		os.Exit(1)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		os.Exit(1)
	}

	// Load the public key from the file
	publicKeyFile, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		fmt.Println("Error loading public key file:", err)
		os.Exit(1)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyFile)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		os.Exit(1)
	}

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte("super secret message"),
		nil)
	if err != nil {
		panic(err)
	}

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	// We get back the original information in the form of bytes, which we
	// the cast to a string and print
	fmt.Println("decrypted message: ", string(decryptedBytes))
}
