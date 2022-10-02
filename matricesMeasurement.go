package main

import "fmt"

func main() {
	binaryMatrix := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(Area(binaryMatrix))
}

type regionStruct struct {
	column int
	row    int
}

func Area(binaryMatrix [][]int) int {
	columns := len(binaryMatrix)
	rows := len(binaryMatrix[0])
	var region []regionStruct
	var maxRegCount, regCount int

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if binaryMatrix[i][j] == 1 {
				regCount = maxRegCount
				region = append(region, regionStruct{i, j})
				binaryMatrix[i][j] = 4 //for it not to be calculated anymore
				for k := 0; k < len(region); k++ {

					regElem := region[k]
					region = append(region[:k], region[k+1:]...) //delete regElem from region array
					regCount += 1

					//directions
					c, r := regElem.column, regElem.row
					if c+1 < columns && binaryMatrix[c+1][r] == 1 { // top
						region = append(region, regionStruct{c + 1, r})
						binaryMatrix[c+1][r] = 4
					}
					if c-1 >= 0 && binaryMatrix[c-1][r] == 1 { // bottom
						region = append(region, regionStruct{c - 1, r})
						binaryMatrix[c-1][r] = 4
					}
					if r+1 < rows && binaryMatrix[c][r+1] == 1 { // right
						region = append(region, regionStruct{c, r + 1})
						binaryMatrix[c][r+1] = 4
					}
					if r-1 >= 0 && binaryMatrix[c][r-1] == 1 { // left
						region = append(region, regionStruct{c, r - 1})
						binaryMatrix[c][r-1] = 4
					}
					if c+1 < columns && r+1 < rows && binaryMatrix[c+1][r+1] == 1 { // right top diagonal
						region = append(region, regionStruct{c + 1, r + 1})
						binaryMatrix[c+1][r+1] = 4
					}
					if c-1 >= 0 && r+1 < rows && binaryMatrix[c-1][r+1] == 1 { // right bottom diagonal
						region = append(region, regionStruct{c - 1, r + 1})
						binaryMatrix[c-1][r+1] = 4
					}
					if c-1 >= 0 && r-1 >= 0 && binaryMatrix[c-1][r-1] == 1 { //left bottom diagonal
						region = append(region, regionStruct{c - 1, r - 1})
						binaryMatrix[c-1][r-1] = 4
					}
					if c+1 < columns && r-1 >= 0 && binaryMatrix[c+1][r-1] == 1 { // right bottom diagonal
						region = append(region, regionStruct{c + 1, r - 1})
						binaryMatrix[c+1][r-1] = 4
					}
				}
				if maxRegCount < regCount {
					maxRegCount = regCount
				}
				regCount = 0
			}
		}
	}
	return maxRegCount
}
