package static

import (
	"homework/pkg/vm/value"
	"reflect"
	"unsafe"
)

func NewFloats(vs []float64) *Floats {
	return &Floats{
		xs: vs,
	}
}

func (a *Floats) Size() int {
	return len(a.xs) * 8
}

func (a *Floats) Count() int {
	return len(a.xs)
}

func (a *Floats) Show() ([]byte, error) {
	hp := *(*reflect.SliceHeader)(unsafe.Pointer(&a.xs))
	hp.Len *= 8
	hp.Cap *= 8
	return *(*[]byte)(unsafe.Pointer(&hp)), nil
}

func (a *Floats) Read(cnt int, data []byte) error {
	hp := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	hp.Len = cnt
	hp.Cap = cnt
	a.xs = *(*[]float64)(unsafe.Pointer(&hp))
	return nil
}

func (a *Floats) Append(v interface{}) error {
	a.xs = append(a.xs, v.([]float64)...)
	return nil
}

func (a *Floats) ValueList() []value.Value {
	vs := make([]value.Value, len(a.xs))
	for i, j := 0, len(a.xs); i < j; i++ {
		vs[i] = value.NewFloat(a.xs[i])
	}
	return vs
}
