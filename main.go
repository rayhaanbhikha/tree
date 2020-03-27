package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	pipe         = "│"
	treeNode     = "├──"
	lastTreeNode = "└──"
	separtor     = "   "
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dir := os.Args[1]
	tree(0, dir)

}

func tree(indent int, dir string) {
	files, err := ioutil.ReadDir(dir)
	checkErr(err)
	printLeafNode(indent, path.Base(dir))
	for _, file := range files {

		fileName := file.Name()
		//  rules to ignore
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}

		if file.IsDir() {
			tree(indent+1, path.Join(dir, fileName))
		} else {
			printLeafNode(indent+1, fileName)
		}
	}
}

func printLeafNode(indent int, fileName string) {
	fmt.Println(indent, fileName)
}
