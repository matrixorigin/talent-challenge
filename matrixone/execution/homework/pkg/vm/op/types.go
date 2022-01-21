package op

import (
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/value"
)

const (
	Relation = iota
	Projection
)

type OP interface {
	Name() string
	Attributes() []metadata.Attribute
	Read([]string) (map[string][]value.Value, error)
}
