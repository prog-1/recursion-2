package main

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	for _, tc := range []struct {
		input [][]int
		want  int
	}{
		{[][]int{
			{0, 0},
			{0, 0},
		}, 0},
		{[][]int{
			{1, 0},
			{0, 0},
		}, 1},
		{[][]int{
			{1, 0},
			{0, 1},
		}, 2},
		{[][]int{
			{1, 1, 1},
			{0, 0, 1},
		}, 4},
		{[][]int{
			{1, 1, 1},
			{0, 1, 1},
		}, 5},
		{[][]int{
			{1, 1, 1},
			{1, 1, 1},
		}, 6},
		{[][]int{
			{0, 0, 1, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}, 6},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			s := make([][]int, len(tc.input)) // because the function Area() destroys the original 2d slice
			for i := range tc.input {
				s[i] = make([]int, len(tc.input[0]))
				copy(s[i], tc.input[i])
			}
			if got := Area(s); got != tc.want {
				t.Errorf("Area(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
