package main

import "fmt"

func Area(a, b int, m [][]int) (res int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if m[i][j] == 1 {
				res += Max(res, FloodFill(a, b, i, j, m))
			}
		}
	}
	return res
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func FloodFill(a, b, i, j int, m [][]int) int {
	if i >= a || j >= b || i < 0 || j < 0 || m[i][j] < 1 {
		return 0
	}
	m[i][j] = 0
	return 1 +
		FloodFill(a, b, i+1, j, m) + //right
		FloodFill(a, b, i-1, j, m) + //left
		FloodFill(a, b, i, j+1, m) + //down
		FloodFill(a, b, i, j-1, m) + // up
		FloodFill(a, b, i+1, j-1, m) + // diagonally up right
		FloodFill(a, b, i-1, j-1, m) + // diagonally up left
		FloodFill(a, b, i+1, j+1, m) + // diagonally down right
		FloodFill(a, b, i-1, j+1, m) // diagonally down left
}

func main() {
	a, b := 4, 5
	m := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(Area(a, b, m))
}
