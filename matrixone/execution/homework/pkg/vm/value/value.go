package value

import (
	"fmt"
	"homework/pkg/vm/types"
	"strings"
)

func Compare(a, b Value) int {
	if x, y := a.ResolvedType().Oid, b.ResolvedType().Oid; x == y || isNum(x, y) {
		return a.Compare(b)
	}
	panic(makeUnsupportedComparisonMessage(a, b))
}

func isNum(x, y uint32) bool {
	switch {
	case x == types.T_int && isInt(y):
		return true
	case y == types.T_int && isInt(x):
		return true
	case x == types.T_float && isFloat(y):
		return true
	case y == types.T_float && isFloat(x):
		return true
	}
	return false
}

func isInt(typ uint32) bool {
	return typ == types.T_int
}

func isFloat(typ uint32) bool {
	return typ == types.T_float
}

// makeParseError returns a parse error using the provided string and type. An
// optional error can be provided, which will be appended to the end of the
// error string.
func makeParseError(s string, typ *types.T, err error) error {
	if err != nil {
		return fmt.Errorf("could not parse %q as type %s: %v", s, typ, err)
	}
	return fmt.Errorf("could not parse %q as type %s", s, typ)
}

func makeUnsupportedComparisonMessage(d1, d2 Value) error {
	return fmt.Errorf("unsupported comparison: %s to %s", d1.ResolvedType(), d2.ResolvedType())
}

func isCaseInsensitivePrefix(prefix, s string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return strings.EqualFold(prefix, s[:len(prefix)])
}
