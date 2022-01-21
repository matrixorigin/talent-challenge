package extend

import (
	"fmt"
	"homework/pkg/vm/extend/overload"
	"homework/pkg/vm/types"
	"homework/pkg/vm/util"
	"homework/pkg/vm/value"
)

func (e *BinaryExtend) Attributes() []string {
	return util.MergeAttributes(e.Left.Attributes(), e.Right.Attributes())
}

func (e *BinaryExtend) ReturnType() uint32 {
	switch e.Op {
	case overload.Plus:
		lt, rt := e.Left.ReturnType(), e.Right.ReturnType()
		return returnType(lt, rt)
	case overload.Minus:
		lt, rt := e.Left.ReturnType(), e.Right.ReturnType()
		return returnType(lt, rt)
	}
	return 0
}

func (e *BinaryExtend) Eval(mp map[string][]value.Value) ([]value.Value, uint32, error) {
	l, lt, err := e.Left.Eval(mp)
	if err != nil {
		return nil, 0, err
	}
	r, rt, err := e.Right.Eval(mp)
	if err != nil {
		return nil, 0, err
	}
	return overload.BinaryEval(e.Op, lt, rt, l, r)
}

func (e *BinaryExtend) String() string {
	switch e.Op {
	case overload.Plus:
		return fmt.Sprintf("%s + %s", e.Left.String(), e.Right.String())
	case overload.Minus:
		return fmt.Sprintf("%s - %s", e.Left.String(), e.Right.String())
	}
	return ""
}

func (e *ParenExtend) Attributes() []string {
	return e.E.Attributes()
}

func (e *ParenExtend) Eval(mp map[string][]value.Value) ([]value.Value, uint32, error) {
	return e.E.Eval(mp)
}

func (e *ParenExtend) String() string {
	return "(" + e.E.String() + ")"
}

func (a *Attribute) Attributes() []string {
	return []string{a.Name}
}

func (a *Attribute) ReturnType() uint32 {
	return a.Type
}

func (a *Attribute) Eval(mp map[string][]value.Value) ([]value.Value, uint32, error) {
	if vs, ok := mp[a.Name]; ok {
		return vs, a.Type, nil
	}
	return nil, 0, fmt.Errorf("attribute '%s' not exist", a.Name)
}

func (a *Attribute) String() string {
	return a.Name
}

func returnType(x, y uint32) uint32 {
	if x == y {
		return x
	}
	if x == types.T_int || x == types.T_float {
		return y
	}
	return x
}
