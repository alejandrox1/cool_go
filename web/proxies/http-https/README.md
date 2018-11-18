# [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c)

> Handling of HTTP is a matter of parsing request, passing such request to 
> destination server, reading response and passing it back to the client. 
> All we need for that is built-in HTTP server and client (net/http). 
> HTTPS is different as it’ll use technique called `HTTP CONNECT` tunneling. 
> First client sends request using `HTTP CONNECT` method to set up the tunnel
> between the client and destination server. When such tunnel consisting of
> two TCP connections is ready, client starts regular TLS handshake with 
> destination server to establish secure connection and later send requests
> and receive responses.

## HTTP CONNECT Tunneling

> Suppose client wants to use either HTTPS or WebSockets in order to talk to 
> server. Client is aware of using proxy. Simple HTTP request / response flow
> cannot be used since client needs to e.g. establish secure connection with 
> server (HTTPS) or wants to use other protocol over TCP connection 
> (WebSockets). Technique which works is to use 
> [`HTTP CONNECT` method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT). 
> It tells the proxy server to establish TCP connection with destination server
> and when done to proxy the TCP stream to and from the client. This way proxy
> server won’t terminate SSL but will simply pass data between client and
> destination server so these two parties can establish secure connection.

> Presented code is not a production-grade solution. It lacks e.g. 
> [`handling hop-by-hop` headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#hbh),
> setting up timeouts while copying data between two connections or the ones
> exposed by net/http — more on this in 
> [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/).
