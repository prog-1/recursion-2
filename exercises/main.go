package main

import "fmt"

// Calculates all ones in the array
func area(s [][]int) int {
	n := len(s)
	m := len(s[0])
	checked := make([][]bool, n)
	for i := range checked {
		checked[i] = make([]bool, m)
	}
	var areas []int
	for i := range s {
		for j := range s[i] {
			areas = append(areas, dfs(s, i, j, checked))
		}
	}
	max := areas[0]
	for _, area := range areas {
		if area > max {
			max = area
		}
	}
	return max
}

func dfs(s [][]int, i, j int, checked [][]bool) int {
	if i < 0 || i > len(s)-1 || j < 0 || j > len(s[0])-1 || checked[i][j] || s[i][j] == 0 {
		return 0
	}
	checked[i][j] = true

	return s[i][j] + dfs(s, i, j-1, checked) + dfs(s, i-1, j-1, checked) + dfs(s, i-1, j, checked) + dfs(s, i-1, j+1, checked) + dfs(s, i, j+1, checked) + dfs(s, i+1, j+1, checked) + dfs(s, i+1, j, checked) + dfs(s, i+1, j-1, checked)
}

func main() {
	s := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1},
	}

	fmt.Println(area(s))
}
