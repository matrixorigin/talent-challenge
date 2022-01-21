package value

import "homework/pkg/vm/types"

func (a *Int) Eval(mp map[string][]Value) ([]Value, uint32, error) {
	rs := make([]Value, length(mp))
	for i := range rs {
		rs[i] = a
	}
	return rs, types.T_int, nil
}

func (a *Float) Eval(mp map[string][]Value) ([]Value, uint32, error) {
	rs := make([]Value, length(mp))
	for i := range rs {
		rs[i] = a
	}
	return rs, types.T_float, nil
}

func length(mp map[string][]Value) int {
	for _, v := range mp {
		return len(v)
	}
	return 0
}
