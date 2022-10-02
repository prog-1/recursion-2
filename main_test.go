package main

import (
	"reflect"
	"testing"
)

func TestListFiles(t *testing.T) {
	want := []string{"test/README.md", "test/ds/main.go", "test/ds/nm.md", "test/go.mod", "test/main.go"}
	got := ListFiles("./test")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestArea(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input [][]int
		want  int
	}{
		{"one area", [][]int{
			[]int{0, 0, 1, 1, 0},
			[]int{1, 0, 1, 1, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
		}, 6},
		{"two areas", [][]int{
			[]int{0, 0, 1, 1, 0},
			[]int{1, 0, 1, 1, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 0, 0, 0, 1},
		}, 6},
		{"one small area", [][]int{
			[]int{1, 1, 1},
			[]int{0, 0, 1},
		}, 4},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Area(tc.input); got != tc.want {
				t.Errorf("wrong answer in test case named \"%v\", got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}
