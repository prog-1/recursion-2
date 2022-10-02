package main

import "testing"

func TestArea(t *testing.T) {
	for _, tc := range []struct {
		a    [][]int
		n, m int
		res  int
	}{
		{[][]int{
			{0, 0, 1, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}, 6, 5, 9},
		{[][]int{
			{0, 0, 1, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}, 4, 5, 6},
		{[][]int{
			{1, 1, 1},
			{0, 0, 1},
		}, 2, 3, 4},
		{[][]int{
			{0, 0},
			{0, 0},
		}, 2, 2, 0},
		{[][]int{
			{1},
		}, 1, 1, 1},
		{[][]int{
			{1},
			{0},
		}, 2, 1, 1},
	} {
		if res1 := Area(tc.n, tc.m, tc.a); res1 != tc.res {
			t.Errorf("got =%v, want = %v", res1, tc.res)
		}
	}
}
