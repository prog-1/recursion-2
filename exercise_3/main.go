package main

import "fmt"

func Area(m [][]int) (max int) {
	for row := range m {
		for col := range m[row] {
			result := Island(m, row, col)
			if result > max {
				max = result
			}
		}
	}
	return
}

func Island(m [][]int, row, col int) int {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) || m[row][col] == 0 {
		return 0
	}
	m[row][col] = 0
	return 1 +
		Island(m, row-1, col) + // North
		Island(m, row, col-1) + // West
		Island(m, row, col+1) + // East
		Island(m, row+1, col) + // South
		Island(m, row-1, col-1) +
		Island(m, row-1, col+1) +
		Island(m, row+1, col-1) +
		Island(m, row+1, col+1)
}

func main() {
	m := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(Area(m))
}
