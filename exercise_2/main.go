package main

import (
	"fmt"
	"os"
)

func ListFiles(dir string) []string {
	var result []string
	directories, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, directory := range directories {
		if directory.IsDir() {
			result = append(result, ListFiles(dir+"/"+directory.Name())...)
		} else { // is file
			result = append(result, dir+"/"+directory.Name())
		}
	}
	return result
}

func main() {
	s := ListFiles("exercise_2/testdata")
	for _, i := range s {
		fmt.Println(i)
	}
}
