package main

import (
	"reflect"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Run("File test", func(t *testing.T) {
		want := []string{"files.go", "matrix.go", "recursion2_test.go"}
		got := ListFiles(".")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}

func TestMatrix(t *testing.T) {
	t.Run("Cell test", func(t *testing.T) {
		input := [][]int{
			{0, 0, 1, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		want := 6
		got := Area(input)
		if got != want {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}
