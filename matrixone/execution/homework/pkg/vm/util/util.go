package util

import "fmt"

func Contain(xs, ys []string) error {
	mp := make(map[string]struct{})
	for _, y := range ys {
		mp[y] = struct{}{}
	}
	for _, x := range xs {
		if _, ok := mp[x]; !ok {
			return fmt.Errorf("'%s' not in '%v'", x, ys)
		}
	}
	return nil
}

func MergeAttributes(xs, ys []string) []string {
	var rs []string

	mp := make(map[string]struct{})
	for i, j := 0, len(xs); i < j; i++ {
		if _, ok := mp[xs[i]]; !ok {
			mp[xs[i]] = struct{}{}
			rs = append(rs, xs[i])
		}
	}
	for i, j := 0, len(ys); i < j; i++ {
		if _, ok := mp[ys[i]]; !ok {
			mp[ys[i]] = struct{}{}
			rs = append(rs, ys[i])
		}
	}
	return rs
}
