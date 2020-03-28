package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/rayhaanbhikha/tree/build"
)

const (
	pipe         = "│   "
	separtor     = "    "
	treeNode     = "├──"
	lastTreeNode = "└──"
)

var (
	dir      string
	maxLevel int
)

func init() {
	flag.StringVar(&dir, "d", ".", "directory to list tree")
	flag.IntVar(&maxLevel, "l", 5, "max level")
	flag.Parse()
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("tree version %s\n", build.Version)
		return
	}

	fmt.Println(dir)
	err := tree("", dir, 0)
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
