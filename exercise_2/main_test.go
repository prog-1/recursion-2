package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListFiles(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  []string
	}{
		{"testdata",
			[]string{"testdata/home/yarcat/README.md",
				"testdata/usr/bin/cat",
				"testdata/usr/bin/gcc",
				"testdata/usr/bin/ls",
				"testdata/usr/local/bin/go1.19",
				"testdata/usr/local/bin/mytool"}},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := ListFiles(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ListFiles(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
