package value

import (
	"fmt"
	"homework/pkg/vm/types"
	"strconv"
)

func NewInt(v int64) *Int {
	r := Int(v)
	return &r
}

func AsInt(v interface{}) (Int, bool) {
	switch t := v.(type) {
	case *Int:
		return *t, true
	default:
		return 0, false
	}
}

// MustBeInt attempts to retrieve a Int from a value, panicking if the
// assertion fails.
func MustBeInt(v interface{}) int64 {
	i, ok := AsInt(v)
	if !ok {
		panic(fmt.Errorf("expected *Int, found %T", v))
	}
	return int64(i)
}

func GetInt(v Value) (Int, error) {
	if i, ok := v.(*Int); ok {
		return *i, nil
	}
	return 0, fmt.Errorf("cannot convert %s to type %s", v.ResolvedType(), types.Int)
}

func (a *Int) String() string {
	return strconv.FormatInt(int64(*a), 10)
}

func (_ *Int) ResolvedType() *types.T {
	return types.Int
}

// ParseInt parses and returns the *Int value represented by the provided
// string, or an error if parsing is unsuccessful.
func ParseInt(s string) (*Int, error) {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return nil, makeParseError(s, types.Int, err)
	}
	return NewInt(i), nil
}

func (a *Int) Compare(v Value) int {
	var x, y int64

	x = int64(*a)
	switch b := v.(type) {
	case *Int:
		y = int64(*b)
	default:
		panic(makeUnsupportedComparisonMessage(a, v))
	}
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	default:
		return 0
	}
}

func (_ *Int) Size() int            { return 9 }
func (_ *Int) Attributes() []string { return []string{} }
