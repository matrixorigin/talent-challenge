package overload

import (
	"fmt"
	"homework/pkg/vm/types"
	"homework/pkg/vm/value"
)

func BinaryEval(op int, ltyp, rtyp uint32, as, bs []value.Value) ([]value.Value, uint32, error) {
	if os, ok := BinOps[op]; ok {
		for _, o := range os {
			if binaryCheck(op, o.LeftType, o.RightType, ltyp, rtyp) {
				rs, err := o.Fn(as, bs)
				return rs, o.ReturnType, err
			}
		}
	}
	return nil, 0, fmt.Errorf("%s not yet implemented for %s, %s", OpName[op], &types.T{ltyp}, &types.T{rtyp})
}

func binaryCheck(op int, arg0, arg1 uint32, val0, val1 uint32) bool {
	return arg0 == val0 && arg1 == val1
}

// BinOps contains the binary operations indexed by operation type.
var BinOps = map[int][]*BinOp{
	Plus: {
		&BinOp{
			LeftType:   types.T_int,
			RightType:  types.T_int,
			ReturnType: types.T_int,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewInt(value.MustBeInt(as[i]) + value.MustBeInt(bs[i]))
				}
				return rs, nil
			},
		},
		&BinOp{
			LeftType:   types.T_int,
			RightType:  types.T_float,
			ReturnType: types.T_float,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewFloat(float64(value.MustBeInt(as[i])) + value.MustBeFloat(bs[i]))
				}
				return rs, nil
			},
		},
		&BinOp{
			LeftType:   types.T_float,
			RightType:  types.T_int,
			ReturnType: types.T_float,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewFloat(value.MustBeFloat(as[i]) + float64(value.MustBeInt(bs[i])))
				}
				return rs, nil
			},
		},
	},
	Minus: {
		&BinOp{
			LeftType:   types.T_int,
			RightType:  types.T_int,
			ReturnType: types.T_int,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewInt(value.MustBeInt(as[i]) - value.MustBeInt(bs[i]))
				}
				return rs, nil
			},
		},
		&BinOp{
			LeftType:   types.T_int,
			RightType:  types.T_float,
			ReturnType: types.T_float,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewFloat(float64(value.MustBeInt(as[i])) - value.MustBeFloat(bs[i]))
				}
				return rs, nil
			},
		},
		&BinOp{
			LeftType:   types.T_float,
			RightType:  types.T_int,
			ReturnType: types.T_float,
			Fn: func(as, bs []value.Value) ([]value.Value, error) {
				rs := make([]value.Value, len(as))
				for i := range as {
					rs[i] = value.NewFloat(value.MustBeFloat(as[i]) - float64(value.MustBeInt(bs[i])))
				}
				return rs, nil
			},
		},
	},
}
