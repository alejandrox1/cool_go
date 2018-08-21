package tracer

import (
    "fmt"
    "io"
)


// Tracer is the interface that describes an object capable of tracing events
// throughout a service.
type Tracer interface {
    Trace(...interface{})
}


// tracer implements the Tracer interface by writing to an io.Writer.
type tracer struct {
    out io.Writer
}


// New creates a new Tracer that will write the output to the specified
// io.Writer.
func New(w io.Writer) Tracer {
    return &tracer{out: w}
}

// Trace writes the arguments to this tracer to the io.Writer.
func (t *tracer) Trace (a ...interface{}) {
    fmt.Fprint(t.out, a...)
    fmt.Fprintln(t.out)
}


// nilTracer 
type nilTracer struct {}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
    return &nilTracer{}
}
