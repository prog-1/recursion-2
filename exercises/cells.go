package main

// Returns area(cell count) of the largest region
func Area(A [][]int) int {
	// Creating veriables:
	n, m := len(A), len(A[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	var areas []int

	// Getting all region areas:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				if !visited[i][j] {
					var area int
					DFS(i, j, A, &visited, &area)
					areas = append(areas, area)
				}
			}
		}
	}

	// Returning greatest area:
	return Max(areas)
}

func DFS(i, j int, A [][]int, visited *[][]bool, area *int) {
	(*visited)[i][j] = true
	*area++

	n, m := len(A), len(A[0])
	if i-1 >= 0 && A[i-1][j] == 1 && !(*visited)[i-1][j] { // Up
		DFS(i-1, j, A, visited, area)
	}
	if j+1 < m && A[i][j+1] == 1 && !(*visited)[i][j+1] { // Right
		DFS(i, j+1, A, visited, area)
	}
	if i+1 < n && A[i+1][j] == 1 && !(*visited)[i+1][j] { // Down
		DFS(i+1, j, A, visited, area)
	}
	if j-1 > 0 && A[i][j-1] == 1 && !(*visited)[i][j-1] { // Left
		DFS(i, j-1, A, visited, area)
	}
	if i-1 >= 0 && j-1 >= 0 && A[i-1][j-1] == 1 && !(*visited)[i-1][j-1] { // LeftUp
		DFS(i-1, j-1, A, visited, area)
	}
	if i-1 >= 0 && j+1 < m && A[i-1][j+1] == 1 && !(*visited)[i-1][j+1] { // RightUp
		DFS(i-1, j+1, A, visited, area)
	}
	if i+1 < n && j+1 < m && A[i+1][j+1] == 1 && !(*visited)[i+1][j+1] { // RightDown
		DFS(i+1, j+1, A, visited, area)
	}
	if i+1 < n && j-1 >= 0 && A[i+1][j-1] == 1 && !(*visited)[i+1][j-1] { // LeftDown
		DFS(i+1, j-1, A, visited, area)
	}
}
func Max(s []int) int {
	max := s[0]

	for _, el := range s {
		if el > max {
			max = el
		}
	}

	return max
}

// func main() {
// 	A := [][]int{
// 		{1, 0, 0, 0, 0},
// 		{0, 1, 1, 0, 0},
// 		{0, 1, 0, 1, 0},
// 		{0, 0, 0, 0, 1},
// 	}
// 	fmt.Println(Area(A))
// }
