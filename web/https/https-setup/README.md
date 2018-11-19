# [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)

## TLS Configuration
```go
// Load certificates.
cer, err := tls.LoadX509KeyPair("server.crt", "server.key")

&tls.Config{
    Certificates: []tls.Certificate{cer},
    // Causes servers to use Go's default ciphersuite preferences,
	// which are tuned to avoid attacks. Does nothing on clients.
	PreferServerCipherSuites: true,
	// Only use curves which have assembly implementations.
	CurvePreferences: []tls.CurveID{
		tls.CurveP256,
		tls.X25519, // Go 1.8 only
	},
    MinVersion: tls.VersionTLS12,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
    },
}
```

For a perfct score, you might want to try the following curves:
```go
CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
```
The first two options will however increase the time for a request to be
served.

## HTTP Redirect
```go
srv := &http.Server{
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 5 * time.Second,
	Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Connection", "close")
		url := "https://" + req.Host + req.URL.String()
		http.Redirect(w, req, url, http.StatusMovedPermanently)
	}),
}
go func() { log.Fatal(srv.ListenAndServe()) }()
```

## HTTPS Configuration
```go
srv := &http.Server{
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
    TLSConfig:    tlsConfig,
    Handler:      serveMux,
}
log.Println(srv.ListenAndServeTLS("", ""))
```

## HSTS

If you want to add [HTTP Strict Security Transport](https://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security) you can do as such:
```go
w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
```

The HSTS policy specifies the time which a user agent should only request the
server using a secure connection.
**The HSTS header is only recognized when when sent over an HTTPS connection**.

## References 

* [Simple Golang HTTPS/TLS Examples](https://github.com/denji/golang-tls)
