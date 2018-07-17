package main

import (
    "log"
    "net"
    "net/rpc"

    "testsing_grounds/rpc_tests/rpc_basics/shared"
)


func registerArith(server *rpc.Server, arith shared.Arith) {
    // Registers Arith interface by the name "Arithmetic."
    // To have the name be the same as the object use server.Register.
    server.RegisterName("Arithmetic", arith)
}


func main() {
    //Interface for airthmetic ops.
    arith := new(Arith)

    // Register a new RPC server.
    server := rpc.NewServer()
    registerArith(server, arith)

    ln, err := net.Listen("tcp", ":1234")
    if err != err {
        log.Fatal("listen error: ", err)
    }

    server.Accept(ln)
}
