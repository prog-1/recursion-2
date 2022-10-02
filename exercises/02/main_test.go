package main

import (
	"reflect"
	"testing"
)

func TestListFiles(t *testing.T) {
	want := []string{"exercises/01/main.go", "exercises/02/main.go", "exercises/02/main_test.go", "exercises/03_matrix/main.go", "exercises/03_matrix/main_test.go"}
	got := ListFiles("./exercises")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got = %v want = %v", got, want)
	}
}
