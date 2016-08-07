package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const TOP_FILES int = 10

var fileInfos FileInfos

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfos.Add(FileInfo{path, f.Size()})
	}
	return nil
}

func main() {
	fileInfos = make(FileInfos, TOP_FILES)
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println("collecting infos ...")
	filepath.Walk(root, visit)
	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileInfos)
	fmt.Printf("directory-hotspots:\n  <not yet implemented>\n\n")
	fmt.Printf("tree-hotspots:\n  <not yet implemented>\n\n")
}
