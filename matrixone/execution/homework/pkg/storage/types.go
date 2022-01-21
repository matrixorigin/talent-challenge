package storage

import (
	"homework/pkg/storage/attribute"
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/value"
)

type Relation interface {
	Segments() int

	Attributes() []metadata.Attribute

	AddTuples([]interface{}) error

	Read(int, []string) ([][]value.Value, error)
}

type relation struct {
	mp map[string]int // attribute's name -> index
	md metadata.Metadata
	as []attribute.Attribute
}
