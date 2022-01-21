package types

import "fmt"

const (
	T_null = iota
	T_int  // int64
	T_bool
	T_float // float64
	T_array
	T_string
	T_timestamp
	T_uint8
	T_uint16
	T_uint32
	T_uint64
	T_int8
	T_int16
	T_int32
	T_int64
	T_float32
	T_float64
)

type T struct {
	Oid uint32
}

var (
	Null = &T{T_null}

	Bool = &T{T_bool}

	Array = &T{T_array}

	String = &T{T_string}

	Int = &T{T_int}

	Timestamp = &T{T_timestamp}

	Float = &T{T_float}

	Uint8 = &T{T_uint8}

	Uint16 = &T{T_uint16}

	Uint32 = &T{T_uint32}

	Uint64 = &T{T_uint64}

	Int8 = &T{T_int8}

	Int16 = &T{T_int16}

	Int32 = &T{T_int32}

	Int64 = &T{T_int64}

	Float32 = &T{T_float32}

	Float64 = &T{T_float64}
)

func (t *T) String() string { return t.SQLString() }

func (t *T) SQLString() string {
	switch t.Oid {
	case T_int:
		return "INT"
	case T_null:
		return "NULL"
	case T_bool:
		return "BOOL"
	case T_timestamp:
		return "TIMESTAMP"
	case T_float:
		return "FLOAT"
	case T_array:
		return "ARRAY"
	case T_string:
		return "STRING"
	case T_uint8:
		return "UINT8"
	case T_uint16:
		return "UINT16"
	case T_uint32:
		return "UINT32"
	case T_uint64:
		return "UINT64"
	case T_int8:
		return "INT8"
	case T_int16:
		return "INT16"
	case T_int32:
		return "INT32"
	case T_int64:
		return "INT64"
	case T_float32:
		return "FLOAT32"
	case T_float64:
		return "FLOAT64"
	}
	panic(fmt.Errorf("unexpected oid: %v", t.Oid))
}
