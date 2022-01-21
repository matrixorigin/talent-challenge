package projection

import (
	"errors"
	"fmt"
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/extend"
	"homework/pkg/vm/op"
	"homework/pkg/vm/util"
	"homework/pkg/vm/value"
)

func New(prev op.OP, es []*Extend) *projection {
	return &projection{false, prev, es}
}

func (n *projection) Name() string {
	return n.prev.Name()
}

func (n *projection) Attributes() []metadata.Attribute {
	ns := aliasList(n.es)
	rs := make([]metadata.Attribute, len(n.es))
	for i, e := range n.es {
		rs[i] = metadata.Attribute{
			Name: ns[i],
			Type: e.E.ReturnType(),
		}
	}
	return rs
}

func (n *projection) String() string {
	r := fmt.Sprintf("π([")
	for i, e := range n.es {
		switch i {
		case 0:
			if len(e.Alias) == 0 {
				r += fmt.Sprintf("%s", e.E)
			} else {
				r += fmt.Sprintf("%s -> %s", e.E, e.Alias)
			}
		default:
			if len(e.Alias) == 0 {
				r += fmt.Sprintf(", %s", e.E)
			} else {
				r += fmt.Sprintf(", %s -> %s", e.E, e.Alias)
			}
		}
	}
	r += fmt.Sprintf("], %s)", n.prev)
	return r
}

func (n *projection) Read(attrs []string) (map[string][]value.Value, error) {
	var as [][]string

	attrs = util.MergeAttributes(attrs, []string{})
	es := subExtend(n.es, attrs)
	as = append(as, attributeList(es)) // extend's Attributes
	if !n.isCheck {
		if err := n.check(as[0]); err != nil {
			return nil, err
		}
		if err := util.Contain(attrs, aliasList(n.es)); err != nil {
			return nil, err
		}
		n.isCheck = true
	}
	as[0] = attributeList(es)
	mp, err := n.prev.Read(as[0])
	if err != nil {
		return nil, err
	}
	if mp == nil || len(mp[attrs[0]]) == 0 {
		return nil, nil
	}
	rq := make(map[string][]value.Value)
	for _, e := range es {
		vs, _, err := e.E.Eval(mp)
		if err != nil {
			return nil, err
		}
		if t, ok := e.E.(*extend.Attribute); ok && len(e.Alias) == 0 {
			rq[t.Name] = append(rq[t.Name], vs...)
		} else {
			rq[e.Alias] = append(rq[e.Alias], vs...)
		}
	}
	return rq, nil
}

func (n *projection) check(attrs []string) error {
	for _, e := range n.es {
		if len(e.E.Attributes()) <= 0 {
			return errors.New("must act on attributes")
		}
	}
	as := n.prev.Attributes()
	mp := make(map[string]struct{})
	for _, a := range as {
		mp[a.Name] = struct{}{}
	}
	for _, attr := range attrs {
		if _, ok := mp[attr]; !ok {
			return fmt.Errorf("failed to find attribute '%s'", attr)
		}
	}
	return nil
}

func aliasList(es []*Extend) []string {
	var rs []string

	for _, e := range es {
		if t, ok := e.E.(*extend.Attribute); ok && len(e.Alias) == 0 {
			rs = append(rs, t.Name)
		} else {
			rs = append(rs, e.Alias)
		}
	}
	return rs
}

func attributeList(es []*Extend) []string {
	var rs []string

	mp := make(map[string]struct{})
	for _, e := range es {
		as := e.E.Attributes()
		for i, j := 0, len(as); i < j; i++ {
			if _, ok := mp[as[i]]; !ok {
				mp[as[i]] = struct{}{}
				rs = append(rs, as[i])
			}
		}
	}
	return rs
}

func subExtend(es []*Extend, attrs []string) []*Extend {
	var rs []*Extend

	mp := make(map[string]struct{})
	for i, j := 0, len(attrs); i < j; i++ {
		mp[attrs[i]] = struct{}{}
	}
	as := aliasList(es)
	for i, j := 0, len(es); i < j; i++ {
		if _, ok := mp[as[i]]; ok {
			rs = append(rs, es[i])
		}
	}
	return rs
}
