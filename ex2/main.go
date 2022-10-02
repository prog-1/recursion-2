package main

import (
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
		var in int
		res, roots = cycle(res, roots, files, in, roots[i])
		i++
		files = nil
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
func ex2(root string) []string {
	var roots []string
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	var res []string
	res, roots = cycle(res, roots, files, 0, root)
	if len(roots) != 0 {
		res, _ = cycle0(res, roots, 0)
	}
	return res
}
func main() {

}
