// Simple HTTPS client in Go.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTPS server address")
	certFile := flag.String("certfile", "cert.pem", "trusted CA certificate")
	flag.Parse()

	cert, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	//certPool := x509.NewCertPool()
	//if ok := certPool.AppendCertsFromPEM(cert); !ok {
	//	log.Fatalf("unable to parse cert from %s", *certFile)
	//}
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(cert); !ok {
		log.Println("No certs appended, using system certs only")
	}

	// Option 1
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{
	//			//RootCAs: certPool,
	//			RootCAs: rootCAs,
	//		},
	//	},
	//}

	// Option 2
	// dfltTr := http.DefaultTransport.(*http.Transport)
	// tr := &http.Transport{ // copy default parameters
	// 	Proxy:                 dfltTr.Proxy,
	// 	DialContext:           dfltTr.DialContext,
	// 	MaxIdleConns:          dfltTr.MaxIdleConns,
	// 	IdleConnTimeout:       dfltTr.IdleConnTimeout,
	// 	ExpectContinueTimeout: dfltTr.ExpectContinueTimeout,
	// 	TLSHandshakeTimeout:   dfltTr.TLSHandshakeTimeout,
	// 	TLSClientConfig: &tls.Config{
	// 		// InsecureSkipVerify: true,
	// 		RootCAs: rootCAs,
	// 	},
	// }

	// Option 3
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxConnsPerHost = 100
	tr.DisableKeepAlives = true
	tr.DialContext = (&net.Dialer{
		Timeout:       30 * time.Second,
		KeepAlive:     30 * time.Second,
		FallbackDelay: -1,
	}).DialContext
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            rootCAs,
	}
	client := &http.Client{
		Timeout:   time.Duration(100) * time.Second,
		Transport: tr,
	}

	r, err := client.Get("https://" + *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	html, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r.Status)
	fmt.Printf(string(html))
}
