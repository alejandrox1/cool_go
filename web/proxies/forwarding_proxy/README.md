# [How to Use a Forwarding Proxy](https://gianarb.it/blog/golang-forwarding-proxy)

The idea here is to have some number of services sending their responses to the
proxy before reaching the client.

The brains behind a proxy come from [`HTP CONNECT`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT)

> CONNECT method converts the request connection to a transparent TCP/IP 
> tunnel, usually to facilitate SSL-encrypted communication (HTTPS) through an
> unencrypted HTTP proxy.

* `whoyare.go` HTTP server that returns your remote address.
