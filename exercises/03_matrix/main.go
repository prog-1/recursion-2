package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Area(n, m int, a [][]int) (res int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 1 {
				res = Max(res, floodfill(i, j, n, m, a))
			}
		}
	}
	return res
}

func floodfill(i, j, n, m int, a [][]int) int {
	if i < 0 || j < 0 || i >= n || j >= m || a[i][j] < 1 {
		return 0
	}
	a[i][j] = 0
	return 1 + floodfill(i-1, j, n, m, a) + floodfill(i, j-1, n, m, a) +
		+floodfill(i+1, j, n, m, a) + floodfill(i, j+1, n, m, a) +
		+floodfill(i-1, j-1, n, m, a) + floodfill(i-1, j+1, n, m, a) +
		+floodfill(i+1, j+1, n, m, a) + floodfill(i+1, j-1, n, m, a)
}

func main() {
	n, m := 6, 5
	table := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(Area(n, m, table))
}
