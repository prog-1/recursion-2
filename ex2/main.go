package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func cycle0(res, roots []string, i int) ([]string, []string) {

	if i < len(roots) {
		files, err := os.ReadDir(roots[i])
		if err != nil {
			log.Fatal(err)
		}
		res, roots = cycle(res, roots, files, i, roots[i])
		i++
		return cycle0(res, roots, i)
	}
	return res, roots
}
func cycle(res, roots []string, files []fs.DirEntry, i int, root string) ([]string, []string) {

	if i < len(files) {
		if !files[i].IsDir() {
			res = append(res, files[i].Name())
		} else {
			roots = append(roots, root+"/"+files[i].Name())

		}
		i++
		return cycle(res, roots, files, i, root)

	}
	return res, roots
}
func main() {
	var roots []string
	root := "d:/progs"
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	var res []string
	res, roots = cycle(res, roots, files, 0, root)
	if len(roots) != 0 {
		res, roots = cycle0(res, roots, 0)
	}
	for _, v := range roots {
		fmt.Println(v)
	}
}
