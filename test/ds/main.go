package main

import (
	"fmt"
	"os"
)

func ListFiles(dir string) (fileNames []string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, v := range files {
		if v.IsDir() {
			ListFiles(v.Name())
		} else {
			fileNames = append(fileNames, v.Name())
		}
	}
	return
}

func main() {
	fmt.Println(ListFiles("/home/richard/recursion-2/test"))
}
