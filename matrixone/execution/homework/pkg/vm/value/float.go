package value

import (
	"fmt"
	"homework/pkg/vm/types"
	"math"
	"strconv"
)

func NewFloat(v float64) *Float {
	r := Float(v)
	return &r
}

func AsFloat(v interface{}) (Float, bool) {
	switch t := v.(type) {
	case *Float:
		return *t, true
	default:
		return 0.0, false
	}
}

func MustBeFloat(v interface{}) float64 {
	f, ok := AsFloat(v)
	if !ok {
		panic(fmt.Errorf("expected *Float, found %T", v))
	}
	return float64(f)
}

func GetFloat(v Value) (Float, error) {
	if f, ok := v.(*Float); ok {
		return *f, nil
	}
	return 0, fmt.Errorf("cannot convert %s to type %s", v.ResolvedType(), types.Float)
}

func (a *Float) String() string {
	f := float64(*a)
	if _, frac := math.Modf(f); frac == 0 && -1000000 < *a && *a < 1000000 {
		return fmt.Sprintf("%.1f", f)
	} else {
		return fmt.Sprintf("%g", f)
	}
}

func (_ *Float) ResolvedType() *types.T {
	return types.Float
}

// ParseFloat parses and returns the *Float value represented by the provided
// string, or an error if parsing is unsuccessful.
func ParseFloat(s string) (*Float, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, makeParseError(s, types.Float, err)
	}
	return NewFloat(f), nil
}

func (a *Float) Compare(v Value) int {
	var x, y float64

	x = float64(*a)
	switch b := v.(type) {
	case *Float:
		y = float64(*b)
	default:
		panic(makeUnsupportedComparisonMessage(a, v))
	}
	// NaN sorts before non-NaN
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	case x == y:
		return 0
	}
	if math.IsNaN(x) {
		if math.IsNaN(y) {
			return 0
		}
		return -1
	}
	return 1
}

func (_ *Float) Size() int            { return 9 }
func (_ *Float) Attributes() []string { return []string{} }
