package value

import "homework/pkg/vm/types"

func (a *Int) ReturnType() uint32 {
	return types.T_int
}

func (a *Float) ReturnType() uint32 {
	return types.T_float
}
