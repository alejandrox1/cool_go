package main

import (
    "errors"

    "testsing_grounds/rpc_tests/rpc_basics/shared"
)

type Arith int

func (t *Arith) Multiply(args *shared.Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *shared.Args, quo *shared.Quotient) error {
    if args.B == 0 {
        return errors.New("Divide by zero")
    }

    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}
