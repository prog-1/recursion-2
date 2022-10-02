package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func mainFL() {
	fmt.Println(ListFiles("./testdata"))
}

func ListFiles(dir string) []string {
	return printFiles(dir, readDirect(dir))

}

func readDirect(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func printFiles(dir string, files []fs.DirEntry) []string {
	var s []string
	for _, file := range files {
		var dirFiles []string
		if file.IsDir() {
			dirFiles = ListFiles(dir + "/" + file.Name())
			s = append(s, dirFiles...)
		} else {
			s = append(s, dir+"/"+file.Name())
		}
	}
	return s
}

/*

+---- testdata
|      +----- hedgieDir
|      |          hedgie1.txt
|      |          hedgie2.txt
|      +----- hogieDir
|      |      +------ hogie2Dir
|      |                  hedgie4.txt
|      |      		  hedgie3.txt

Output: [./testdata/hedgieDir/hedgie1.txt ./testdata/hedgieDir/hedgie2.txt ./testdata/hogieDir/hedgie3.txt ./testdata/hogieDir/hogie2Dir/hedgie4.txt]

*/

//TODO
// - take the directory ✔
// - cast for directories in directory ✔
// - save file with directory prefix ✔
