package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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

func main() {
	fmt.Println(ListFiles("./test"))
}
