package storage

import (
	"homework/pkg/storage/attribute"
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/value"
)

func New(md metadata.Metadata) *relation {
	mp := make(map[string]int)
	as := make([]attribute.Attribute, len(md.Attrs))
	for i, attr := range md.Attrs {
		mp[attr.Name] = i
		as[i] = attribute.New(attr.Type)
	}
	return &relation{
		md: md,
		mp: mp,
		as: as,
	}
}

func (r *relation) Segments() int {
	return r.md.Num
}

func (r *relation) Attributes() []metadata.Attribute {
	return r.md.Attrs
}

func (r *relation) AddTuples(ts []interface{}) error {
	for i, j := 0, len(ts); i < j; i++ {
		if err := r.as[i].Append(ts[i]); err != nil {
			return err
		}
	}
	r.md.Num++
	return nil
}

func (r *relation) Read(seg int, attrs []string) ([][]value.Value, error) {
	n := len(attrs)
	vs := make([][]value.Value, n)
	as := make([]attribute.Attribute, n)
	{
		for i, attr := range attrs {
			as[i] = r.as[r.mp[attr]]
		}
	}
	for i := 0; i < n; i++ {
		v, err := as[i].Read(seg)
		if err != nil {
			return nil, err
		}
		vs[i] = v
	}
	return vs, nil
}
