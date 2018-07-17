# RPC Basics

An RPC is when a program causes a procedure to be executed in a different
address space. 
This is a form of client-server interaction, typically implemented via a
request-response message-passing system.

Remote Procedure Calls are a form of interprocess communications.
If on the same host machine, they have distinct virtual address spaces.

From here on we wil follow this tutorial: 
[How to build RPC server in golang (step by step with examples)](https://parthdesai.me/articles/2016/05/20/go-rpc-server/)

## Overview
We will build an interface with two methods: `Multiply` and `Divide`.
There will be two structs used to pass arguments from client to server;
`Args` will represent the intput parameters to `Multiply` and `Divide`, 
`Quotient` the output of
`Divide`.


## RPC Server
We can implement an RPC server by listening for incoming connections using the
HTTP protocol and then switching to rpc, which would allow us easily
authenticate RPC clients using commonly used HTTP authentication methods.

Internally, the server would listen for HTTP `CONNECT` method and the uses
`http.Hijacker` to hijack the connection.

The other option is to listen for connections directly.
