package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
)

const TOP_FILES int = 10

var fileHotspots FileInfos
var dirHotspotsCtx *DirHotspotsContext

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfo := FileInfo{path, f.Size()}
		fileHotspots.Add(fileInfo)
		dirHotspotsCtx.AddFile(fileInfo)
	} else {
		dirInfo := &DirInfo{Name: path, Children: []string{}}
		dirHotspotsCtx.AddDir(dirInfo)
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
	root := "."
	dirHotspotsCtx.root = root

	fmt.Println("collecting infos ...")
	filepath.Walk(root, visit)

	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileHotspots)
	fmt.Printf("directory-hotspots:\n%v\n", printDirInfos(dirHotspotsCtx.GetDirHotspots(TOP_FILES), printSize))
	fmt.Printf("tree-hotspots:\n%v\n\n", printDirInfos(dirHotspotsCtx.GetTreeHotspots(TOP_FILES), printTotalSize))
}

type DirPrinter func(info *DirInfo) string

func printDirInfos(infos DirInfos, printer DirPrinter) string {
	result := ""
	for _, info := range infos {
		result += printer(info) + "\n"
	}
	return result
}

func printSize(info *DirInfo) string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.Size)), info.Name)
}

func printTotalSize(info *DirInfo) string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.TotalSize)), info.Name)
}
