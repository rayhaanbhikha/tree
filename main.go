package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
)

const (
	pipe         = "│   "
	separtor     = "    "
	treeNode     = "├──"
	lastTreeNode = "└──"
	maxLevel     = 3
)

var flagDir string

func init() {
	flag.StringVar(&flagDir, "d", ".", "directory to list tree")
	flag.Parse()
}

func main() {
	fmt.Println(flagDir)
	err := tree("", flagDir, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func tree(indent string, dir string, clevel int) error {
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

		if clevel == maxLevel {
			return nil
		}

		newDir := path.Join(dir, fileName)
		if index == n-1 {
			err = tree(indent+separtor, newDir, clevel+1)
		} else {
			err = tree(indent+pipe, newDir, clevel+1)
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
