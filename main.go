package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const TOP_FILES int = 10

var fileHotspots FileInfos
var dirHotspotsCtx *DirHotspotsContext

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfo := FileInfo{path, f.Size()}
		fileHotspots.Add(fileInfo)
		dirHotspotsCtx.Add(fileInfo)
	}
	return nil
}

func main() {
	fileHotspots = make(FileInfos, TOP_FILES)
	dirHotspotsCtx = &DirHotspotsContext{
		dirInfos:   DirInfos{},
		dirInfoIdx: DirInfoIdx{},
	}

	flag.Parse()
	root := flag.Arg(0)

	fmt.Println("collecting infos ...")
	filepath.Walk(root, visit)

	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileHotspots)
	fmt.Printf("directory-hotspots:\n%v\n", dirHotspotsCtx.GetHotspots(TOP_FILES))
	fmt.Printf("tree-hotspots:\n  <not yet implemented>\n\n")
}
