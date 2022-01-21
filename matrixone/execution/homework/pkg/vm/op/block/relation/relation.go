package relation

import (
	"homework/pkg/storage"
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/value"
)

func New(seg int, r storage.Relation) *relation {
	return &relation{
		r:   r,
		seg: seg,
	}
}

func (r *relation) Name() string {
	return "R"
}

func (r *relation) Attributes() []metadata.Attribute {
	return r.r.Attributes()
}

func (r *relation) Read(attrs []string) (map[string][]value.Value, error) {
	if r.r == nil {
		return nil, nil
	}
	defer func() { r.r = nil }()
	vs, err := r.r.Read(r.seg, attrs)
	if err != nil {
		return nil, err
	}
	rp := make(map[string][]value.Value)
	for i, attr := range attrs {
		rp[attr] = vs[i]
	}
	return rp, nil
}
