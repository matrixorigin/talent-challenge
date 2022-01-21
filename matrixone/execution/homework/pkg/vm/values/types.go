package values

import "homework/pkg/vm/value"

type Values interface {
	Size() int
	Count() int

	Read(int, []byte) error
	Show() ([]byte, error)

	Append(interface{}) error

	ValueList() []value.Value
}
