package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"

    "testsing_grounds/rpc_tests/rpc_basics/shared"
)

type Arith struct {
    client *rpc.Client
}

func (t *Arith) Multiply(a, b int) int {
    args := &shared.Args{a, b}
    var reply int
    err := t.client.Call("Arithmetic.Multiply", args, &reply)
    if err != nil {
        log.Fatal("Arithmetic error in Multiply: ", err)
    }
    return reply
}

func (t *Arith) Divide(a, b int) shared.Quotient {
    args := &shared.Args{a, b}
    var reply shared.Quotient
    err := t.client.Call("Arithmetic.Divide", args, &reply)
    if err != nil {
        log.Fatal("Arithmetic error in Divide: ", err)
    }
    return reply
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	arith := &Arith{client: rpc.NewClient(conn)}

	fmt.Println(arith.Multiply(5, 6))
	fmt.Println(arith.Divide(500, 10))
}
