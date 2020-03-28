package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	pipe         = "│   "
	separtor     = "    "
	treeNode     = "├──"
	lastTreeNode = "└──"
)

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	err := tree("", dir)
	if err != nil {
		fmt.Println(err)
	}
}

func tree(indent string, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannot read files from %s: %v", dir, err)
	}
	n := len(files)
	for index, file := range files {

		fileName := file.Name()
		if fileName[0] == '.' || fileName == "node_modules" {
			continue
		}

		print(indent, fileName, index == n-1)
		if !file.IsDir() {
			continue
		}

		newDir := path.Join(dir, fileName)
		if index == n-1 {
			err = tree(indent+separtor, newDir)
		} else {
			err = tree(indent+pipe, newDir)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func print(indent string, fileName string, finalNode bool) {
	if finalNode {
		fmt.Printf("%s%s %s\n", indent, lastTreeNode, fileName)
	} else {
		fmt.Printf("%s%s %s\n", indent, treeNode, fileName)
	}
}
