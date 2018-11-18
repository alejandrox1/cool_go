# [HTTPS for free in Go, with little help of Let's Encrypt](https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html)

Simple web server.

* `-production` - Sets the host policy for the certificate manager to whitelist
your domain. It automatically accepts the terms of service from CA. It also
utilizes a directory to cache certificates and reduce the number of requests.

`-redirect` - it creates an index hanlder that will rewrite the request's
protocol to https keeping host and url the same. Everything else will setup a
server multiplexer just like one would do for any other occasion.
