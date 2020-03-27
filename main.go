package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
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
	d := make([]string, 0)
	fmt.Println(dir)
	tree(d, 0, dir)

}

func tree(d []string, indent int, dir string) {
	files, err := ioutil.ReadDir(dir)
	checkErr(err)
	n := len(files)
	for index, file := range files {

		fileName := file.Name()
		//  rules to ignore
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}

		if file.IsDir() {
			printLeafNode(d, indent+1, fileName, index == n-1)
			tempD := make([]string, 0)
			copy(tempD, d)
			if index == n-1 {
				tempD = append(d, separtor)
			} else {
				tempD = append(d, pipe)
			}
			tree(tempD, indent+1, path.Join(dir, fileName))
		} else {
			printLeafNode(d, indent+1, fileName, index == n-1)
		}
	}
}

func printLeafNode(d []string, indent int, fileName string, finalNode bool) {
	if finalNode {
		fmt.Printf("%s%s %s\n", strings.Join(d, ""), lastTreeNode, fileName)
	} else {
		fmt.Printf("%s%s %s\n", strings.Join(d, ""), treeNode, fileName)
	}
}
