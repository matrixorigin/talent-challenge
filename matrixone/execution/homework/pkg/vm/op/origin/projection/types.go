package projection

import (
	"homework/pkg/vm/extend"
	"homework/pkg/vm/op"
)

type Extend struct {
	Alias string
	E     extend.Extend
}

type projection struct {
	isCheck bool
	prev    op.OP
	es      []*Extend
}
