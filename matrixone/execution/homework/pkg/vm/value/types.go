package value

import "homework/pkg/vm/types"

type Value interface {
	Size() int

	String() string
	Compare(Value) int
	ResolvedType() *types.T

	ReturnType() uint32
	Attributes() []string
	Eval(map[string][]Value) ([]Value, uint32, error)
}

type Int int64
type Float float64
