package static

import (
	"homework/pkg/vm/value"
	"reflect"
	"unsafe"
)

func NewInts(vs []int64) *Ints {
	return &Ints{
		xs: vs,
	}
}

func (a *Ints) Size() int {
	return len(a.xs) * 8
}

func (a *Ints) Count() int {
	return len(a.xs)
}

func (a *Ints) Show() ([]byte, error) {
	hp := *(*reflect.SliceHeader)(unsafe.Pointer(&a.xs))
	hp.Len *= 8
	hp.Cap *= 8
	return *(*[]byte)(unsafe.Pointer(&hp)), nil
}

func (a *Ints) Read(cnt int, data []byte) error {
	hp := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	hp.Len = cnt
	hp.Cap = cnt
	a.xs = *(*[]int64)(unsafe.Pointer(&hp))
	return nil
}

func (a *Ints) Append(v interface{}) error {
	a.xs = append(a.xs, v.([]int64)...)
	return nil
}

func (a *Ints) ValueList() []value.Value {
	vs := make([]value.Value, len(a.xs))
	for i, j := 0, len(a.xs); i < j; i++ {
		vs[i] = value.NewInt(a.xs[i])
	}
	return vs
}
