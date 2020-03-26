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
	tree(0, dir, false)

}

func tree(indent int, dir string, finalDir bool) {
	files, err := ioutil.ReadDir(dir)
	checkErr(err)

	if indent > 0 {
		printLastTreeNode(indent-1, filepath.Base(dir), false)
	}

	dirQueue := make([]string, 0)
	fileQueue := make([]string, 0)

	// filterfiles
	for _, file := range files {

		fileName := file.Name()
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}
		if file.IsDir() {
			dirQueue = append(dirQueue, path.Join(dir, fileName))
		} else {
			fileQueue = append(fileQueue, fileName)
		}
	}

	n := len(fileQueue)
	for index, fileName := range fileQueue {
		if index == n-1 && len(dirQueue) == 0 {
			printLastTreeNode(indent, fileName, finalDir)
		} else {
			printTreeNode(indent, fileName, finalDir)
		}
	}

	dn := len(dirQueue)
	for index, dir := range dirQueue {
		tree(indent+1, dir, index == dn-1)
	}
}

func printLastTreeNode(indent int, fileName string, finalDir bool) {
	s := ""
	for i := 0; i < indent; i++ {
		s += pipe + separtor
	}

	fmt.Printf("%s%s %s\n", s, lastTreeNode, fileName)
}

func printTreeNode(indent int, fileName string, finalDir bool) {
	s := ""
	for i := 0; i < indent; i++ {
		s += pipe + separtor
	}

	fmt.Printf("%s%s %s\n", s, treeNode, fileName)
}
