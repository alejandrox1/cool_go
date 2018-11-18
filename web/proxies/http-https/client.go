package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, err := url.Parse("https://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		// Trust self-signed certificates.
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(u),
		TLSClientConfig: config,
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://googe.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dump)
}
