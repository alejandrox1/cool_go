/*
    Command-line options:
        -production: enables HTTPS on port 443
        -redirect:   redirect HTTP to HTTPS

Taken from https://github.com/kjk/go-cookbook/blob/master/free-ssl-certificates/main.go
*/
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

const (
	htmlIndex = `<html><body>Welcome!</body></html>`
	httpPort  = "127.0.0.1:8080"
)

var (
	flagProduction          = false
	flagRedirectHTTPToHTTPS = false
)

func parseFlags() {
	flag.BoolVar(&flagProduction, "production", false, "Start HTTPS server")
	flag.BoolVar(&flagRedirect, "redirect", false, "Redirect HTTP to HTTPS")
	flag.Parse()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func makeServerFromMux(mux *http.ServeMux) *http.server {
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdelTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", indexHandler)
	return maxkeServerFromMux(mux)
}

func makeHTTPToHTTPSRedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	}

	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleRedirect)
	return makeServerFromMux(mux)
}

func main() {
	parseFlags()

	var m *autocert.Manager
	var httpServer *http.Server

	if flagProduction {
		hostPolicy := func(ctx context.context, host string) error {
			// TODO: change to real host.
			allowedHost := "www.domain.com"
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf("acme/autocert: only %s is allowed", allowedHost)
		}

		dataDir := "."
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(dataDir),
		}

		httpServer = makeHTTPServer()
		httpServer.Addr = ":443"
		httpServer.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", httpServer.Addr)
			if err := httpServer.ListenAndServeTLS("", ""); err != nil {
				log.Fatal(err)
			}
		}()
	}

	if flagRedirectHTTPToHTTPS {
		httpServer = makeHTTPToHTTPSRedirectServer()
	} else {
		httpServer = makeHTTPServer()
	}

	// Allow autocert handle Let's Encrypt callbacks over http.
	if m != nil {
		httpServer.Handler = m.HTTPHandler(httpServer.Handler)
	}

	httpServer.Addr = httpPort
	fmt.Printf("Starting HTTP server on %s\n", httpPort)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
