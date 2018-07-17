package shared

type Arith interface {
    Multiply(args *Args, reply *int) error
    Divide(args *Args, quo *Quotient) error
}
