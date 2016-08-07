package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const TOP_FILES int = 10

var fileHotspots FileInfos
var dirStore DirInfos

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfo := FileInfo{path, f.Size()}
		fileHotspots.Add(fileInfo)
		dirStore.Add(fileInfo)
	}
	return nil
}

func main() {
	fileHotspots = make(FileInfos, TOP_FILES)
	dirStore = make(DirInfos)
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println("collecting infos ...")
	filepath.Walk(root, visit)
	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileHotspots)
	fmt.Printf("directory-hotspots:\n  <not yet implemented>\n\n")
	fmt.Printf("tree-hotspots:\n  <not yet implemented>\n\n")
	fmt.Println("debug:")
	for _, val := range dirStore {
		fmt.Printf("%v\n", val)
	}
}
