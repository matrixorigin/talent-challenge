package extend

import "homework/pkg/vm/value"

type Extend interface {
	String() string
	ReturnType() uint32
	Attributes() []string
	Eval(map[string][]value.Value) ([]value.Value, uint32, error)
}

type BinaryExtend struct {
	Op          int
	Left, Right Extend
}

type ParenExtend struct {
	E Extend
}

type Attribute struct {
	Type uint32
	Name string
}
