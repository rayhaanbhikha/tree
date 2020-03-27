package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	pipe         = "│   "
	treeNode     = "├──"
	lastTreeNode = "└──"
	separtor     = "    "
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	tree("", dir)
}

func tree(indent string, dir string) {
	files, err := ioutil.ReadDir(dir)
	checkErr(err)
	n := len(files)
	for index, file := range files {

		fileName := file.Name()
		//  rules to ignore
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}

		print(indent, fileName, index == n-1)

		if file.IsDir() {
			newfileName := path.Join(dir, fileName)
			if index == n-1 {
				tree(indent+separtor, newfileName)
			} else {
				tree(indent+pipe, newfileName)
			}
		}
	}
}

func print(indent string, fileName string, finalNode bool) {
	if finalNode {
		fmt.Printf("%s%s %s\n", indent, lastTreeNode, fileName)
	} else {
		fmt.Printf("%s%s %s\n", indent, treeNode, fileName)
	}
}
