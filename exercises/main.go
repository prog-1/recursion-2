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
	return dfs(s, 0, 0, checked)
}

func dfs(s [][]int, i, j int, checked [][]bool) int {
	if i < 0 || i > len(s)-1 || j < 0 || j > len(s[0])-1 || checked[i][j] {
		return 0
	}
	checked[i][j] = true

	// dfs(s, i, j-1, result, checked)   // Left
	// dfs(s, i-1, j-1, result, checked) // Top Left
	// dfs(s, i-1, j, result, checked)   // Top
	// dfs(s, i-1, j+1, result, checked) // Top Right
	// dfs(s, i, j+1, result, checked)   // Right
	// dfs(s, i+1, j+1, result, checked) // Down Right
	// dfs(s, i+1, j, result, checked)   // Down
	// dfs(s, i+1, j-1, result, checked) // Down Left

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
