package main

import (
	"errors"
	"fmt"
	"homework/pkg/storage"
	"homework/pkg/storage/metadata"
	"homework/pkg/vm/extend"
	"homework/pkg/vm/extend/overload"
	"homework/pkg/vm/op/block/relation"
	"homework/pkg/vm/op/origin/projection"
	"homework/pkg/vm/types"
	"homework/pkg/vm/value"
	"log"
	"math"
	"time"
)

const (
	Num = 10000000
)

func main() {
	var as []metadata.Attribute
	var ps []*projection.Extend

	{
		as = append(as, metadata.Attribute{types.T_int, "a"})
		as = append(as, metadata.Attribute{types.T_int, "b"})
		as = append(as, metadata.Attribute{types.T_float, "c"})
	}
	r := storage.New(metadata.Metadata{Attrs: as})
	if err := load(r); err != nil {
		log.Fatal(err)
	}
	{
		l := &extend.BinaryExtend{
			Op:    overload.Minus,
			Left:  &extend.Attribute{Name: "a", Type: types.T_int},
			Right: value.NewFloat(99.88),
		}
		r := &extend.BinaryExtend{
			Op:    overload.Plus,
			Left:  &extend.Attribute{Name: "b", Type: types.T_int},
			Right: &extend.Attribute{Name: "c", Type: types.T_float},
		}
		ps = append(ps, &projection.Extend{
			Alias: "A",
			E: &extend.BinaryExtend{
				Op:    overload.Plus,
				Left:  l,
				Right: r,
			},
		})
	}
	o := projection.New(relation.New(0, r), ps) // segments always 0
	t := time.Now()
	for {
		mp, err := o.Read([]string{"A"})
		if err != nil {
			log.Fatal(err)
		}
		if len(mp) == 0 {
			break
		}
		if err := verify(mp["A"]); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("process: %v\n", time.Now().Sub(t))
}

func verify(vs []value.Value) error {
	for i := 0; i < Num; i++ {
		x := float64(i)
		e := (x - 99.88) + (x + 1 + x + 1.1)
		a := value.MustBeFloat(vs[i])
		if math.Abs(e-a) > 0.0001 {
			return errors.New("wrong answer")
		}
	}
	return nil
}

func load(r storage.Relation) error {
	as := make([]int64, Num)
	bs := make([]int64, Num)
	cs := make([]float64, Num)
	for i := 0; i < Num; i++ {
		as[i] = int64(i)
		bs[i] = int64(i) + 1
		cs[i] = float64(i) + 1.1
	}
	return r.AddTuples([]interface{}{as, bs, cs})
}
