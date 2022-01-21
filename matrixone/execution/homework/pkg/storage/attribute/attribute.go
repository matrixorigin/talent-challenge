package attribute

import (
	"homework/pkg/vm/types"
	"homework/pkg/vm/value"
	"homework/pkg/vm/values"
	"homework/pkg/vm/values/static"
)

func New(typ uint32) *attribute {
	return &attribute{typ: typ}
}

func (a *attribute) Read(seg int) ([]value.Value, error) {
	return a.vs[seg].ValueList(), nil
}

func (a *attribute) Append(v interface{}) error {
	var vs values.Values

	vs = newValues(a.typ)
	if err := vs.Append(v); err != nil {
		return err
	}
	a.vs = append(a.vs, vs)
	return nil
}

func newValues(typ uint32) values.Values {
	switch typ {
	case types.T_int:
		return static.NewInts(nil)
	case types.T_float:
		return static.NewFloats(nil)
	}
	return nil
}
