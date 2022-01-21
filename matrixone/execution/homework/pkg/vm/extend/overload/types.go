package overload

import "homework/pkg/vm/value"

const (
	Plus = iota
	Minus
)

var OpName = [...]string{
	Plus:  "+",
	Minus: "-",
}

// UnaryOp is a unary operator.
type UnaryOp struct {
	Typ        uint32
	ReturnType uint32
	Fn         func([]value.Value) ([]value.Value, error)
}

// BinOp is a binary operator.
type BinOp struct {
	LeftType   uint32
	RightType  uint32
	ReturnType uint32

	Fn func([]value.Value, []value.Value) ([]value.Value, error)
}
