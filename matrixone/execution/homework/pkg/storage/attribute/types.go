package attribute

import (
	"homework/pkg/vm/value"
	"homework/pkg/vm/values"
)

type Attribute interface {
	Append(interface{}) error
	Read(int) ([]value.Value, error)
}

type attribute struct {
	typ uint32
	vs  []values.Values
}
