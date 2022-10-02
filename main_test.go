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
