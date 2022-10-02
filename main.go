package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type coord struct {
	x int
	y int
}

func ListFiles(dir string) (fileNames []string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, v := range files {
		if v.IsDir() {
			ListFiles(filepath.Join(dir, v.Name()))
			fileNames = append(fileNames, ListFiles(filepath.Join(dir, v.Name()))...)
		} else {
			fileNames = append(fileNames, filepath.Join(dir, v.Name()))
		}
	}
	return
}

func NotIn(x, y int, s []coord) bool {
	for _, v := range s {
		if v.x == x && v.y == y {
			return false
		}
	}
	return true
}

func findArea(m [][]int, y, x int, checked []coord) (int, []coord) {
	checked = append(checked, coord{x: x, y: y})
	area := 1
	var tmp int
	if y-1 >= 0 {
		if x-1 >= 0 && NotIn(x-1, y-1, checked) && m[y-1][x-1] == 1 {
			tmp, checked = findArea(m, y-1, x-1, checked)
			area += tmp
		}
		if NotIn(x, y-1, checked) && m[y-1][x] == 1 {
			tmp, checked = findArea(m, y-1, x, checked)
			area += tmp
		}
		if x+1 < len(m[y]) && NotIn(x+1, y-1, checked) && m[y-1][x+1] == 1 {
			tmp, checked = findArea(m, y-1, x+1, checked)
			area += tmp
		}
	}
	if y+1 < len(m) {
		if x-1 >= 0 && NotIn(x-1, y+1, checked) && m[y+1][x-1] == 1 {
			tmp, checked = findArea(m, y+1, x-1, checked)
			area += tmp
		}
		if NotIn(x, y+1, checked) && m[y+1][x] == 1 {
			tmp, checked = findArea(m, y+1, x, checked)
			area += tmp
		}
		if x+1 < len(m[y]) && NotIn(x+1, y+1, checked) && m[y+1][x+1] == 1 {
			tmp, checked = findArea(m, y+1, x+1, checked)
			area += tmp
		}
	}
	if x-1 >= 0 && NotIn(x-1, y, checked) && m[y][x-1] == 1 {
		tmp, checked = findArea(m, y, x-1, checked)
		area += tmp
	}
	if x+1 < len(m[y]) && NotIn(x+1, y, checked) && m[y][x+1] == 1 {
		tmp, checked = findArea(m, y, x+1, checked)
		area += tmp
	}
	return area, checked
}

func Area(m [][]int) (area int) {
	var checked []coord
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 1 && NotIn(j, i, checked) {
				var tmp int
				tmp, checked = findArea(m, i, j, checked)
				if tmp > area {
					area = tmp
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(ListFiles("./test"))
	a := [][]int{
		[]int{1, 1, 1},
		[]int{0, 0, 1},
	}
	fmt.Println(Area(a))
}
