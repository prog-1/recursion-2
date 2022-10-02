package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    string
		want []string
	}{
		{"1", "d:/progs/recursion-2/testfolder", []string{"a.txt",
			"m.txt",
			"a.txt",
			"da",
			"t.txt",
			"t1.txt"}},
		{"2", "d:/progs/recursion-2/testfolder/2/3", []string{
			"t.txt"}},
		{"3", "d:/progs/recursion-2/testfolder/2", []string{
			"a.txt",
			"da",
			"t.txt",
			"t1.txt"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := ex2(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
