package main

import (
	"fmt"
	"log"
	"os"
)

func ListFiles(dir string) []string {
	var res []string
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			res = append(res, ListFiles(dir+"/"+file.Name())...)
		} else {
			res = append(res, dir+"/"+file.Name())
		}
	}
	return res
}

func main() {
	fmt.Println(ListFiles("./testfile/03"))
}
