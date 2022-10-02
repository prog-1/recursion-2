package main

import (
	"fmt"
	"log"
	"os"
)

func ListFiles(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())

		var subFiles []string
		if file.IsDir() {
			subFiles = ListFiles(path + "/" + file.Name())
		}
		fileNames = append(fileNames, subFiles...)

	}
	return fileNames
}

func main() {
	files := ListFiles("/files.go")
	fmt.Println(files)
}
