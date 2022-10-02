package main

import (
	"fmt"
	"os"
)

func ListFiles(dir string) []string {
	var files []string
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		panic(err)
	}
	for _, file := range fileInfo {
		files = append(files, dir+"/"+file.Name())
		if file.IsDir() {
			files = append(files, ListFiles(dir+"/"+file.Name())...)
		}
	}
	return files
}

func main() {
	files := ListFiles("./exercises")
	fmt.Println(files)
}
