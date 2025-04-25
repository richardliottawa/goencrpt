package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"reflect"
)

var bundle = []byte(`
-----BEGIN EC PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,99586A658F5D2DAC4A8A3CA387CF71CE

25EtKb7ycOI/5R47fYwpiaNERgYnCxCtcrMXJuOgueuxUXjiU0n93hpUpIQqaTLH
dDKhsR1UHvGJVTV4h577RQ+nEJ5z8K5Y9NWFqzfa/Q5SY43kqqoJ/fS/OCnTmH48
z4bL/dJBDE/a5HwJINgqQhGi9iUkCWUiPQxriJQ0i2s=
-----END EC PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIB2TCCAX+gAwIBAgIUUTZvgwwnbC05WHgIHMXxrbZzr6wwCgYIKoZIzj0EAwIw
QjELMAkGA1UEBhMCWFgxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwT
RGVmYXVsdCBDb21wYW55IEx0ZDAeFw0xOTA1MTQxMzAwMDJaFw0xOTA1MTUxMzAw
MDJaMEIxCzAJBgNVBAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNV
BAoME0RlZmF1bHQgQ29tcGFueSBMdGQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAScgLGx6SXchEo/s0X3AoF0mQkh3bGf9QY0s/2dPqf3/9irwz35DiDGoaP+FDZv
HnUX+D3tUEPhxkLyzWKKT9HHo1MwUTAdBgNVHQ4EFgQU3eB8oRcmvzZrx9Dkb6ma
MMtu1MkwHwYDVR0jBBgwFoAU3eB8oRcmvzZrx9Dkb6maMMtu1MkwDwYDVR0TAQH/
BAUwAwEB/zAKBggqhkjOPQQDAgNIADBFAiAvw/FqAmGbSlBklp6AHJy9kf9VPyhe
RA93ccNQ+7m1fAIhAOXr8c2QsH2oOYRTbn6bPZjkYQ2jLMaxatKhChBIuyZA
-----END CERTIFICATE-----
`)

func main() {
	keyBlock, certsPEM := pem.Decode(bundle)
	fmt.Println(reflect.TypeOf(keyBlock))
	fmt.Println(reflect.TypeOf(certsPEM))
	fmt.Println(x509.IsEncryptedPEMBlock(keyBlock)) // Output: true

	// Decrypt key
	keyDER, err := x509.DecryptPEMBlock(keyBlock, []byte("foobar"))
	if err != nil {
		log.Fatal(err)
	}

	// Update keyBlock with the plaintext bytes and clear the now obsolete
	// headers.
	keyBlock.Bytes = keyDER
	keyBlock.Headers = nil

	// Turn the key back into PEM format so we can leverage tls.X509KeyPair,
	// which will deal with the intricacies of error handling, different key
	// types, certificate chains, etc.
	keyPEM := pem.EncodeToMemory(keyBlock)

	if err := os.WriteFile("decryptedcert.pem", keyPEM, 0644); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote cert.pem\n")

	cert, err := tls.X509KeyPair(certsPEM, keyPEM)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	fmt.Println(reflect.TypeOf(config))
}
