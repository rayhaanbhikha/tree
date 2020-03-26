package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
	n := len(files)
	if indent > 0 {
		printTreeNode(indent-1, filepath.Base(dir))
	}
	// dirQueue := make([]string, 0)
	for index, file := range files {

		fileName := file.Name()
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}
		if file.IsDir() {
			tree(indent+1, path.Join(dir, fileName))
		} else {
			if index == n-1 {
				printLastTreeNode(indent, fileName)
			} else {
				printTreeNode(indent, fileName)
			}
		}
	}
}

func printLastTreeNode(indent int, fileName string) {
	s := ""
	for i := 0; i < indent; i++ {
		s += pipe + separtor
	}
	fmt.Printf("%s%s %s\n", s, lastTreeNode, fileName)
}

func printTreeNode(indent int, fileName string) {
	s := ""
	for i := 0; i < indent; i++ {
		s += pipe + separtor
	}
	fmt.Printf("%s%s %s\n", s, treeNode, fileName)
}
