package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
	"github.com/sgeisbacher/diskusage-analyzer/collector"
	"github.com/sgeisbacher/diskusage-analyzer/context"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

var (
	fileHotspots FileInfos
	ctx          *AnalyzerContext
	dirCollector collector.DirInfoCollector
	topCount     int
)

type DirPrinter func(info *Dir) string

func init() {
	flag.IntVar(&topCount, "n", 10, "limit top-hotspots count")
}

func main() {
	flag.Parse()

	fileHotspots = make(FileInfos, topCount)
	ctx = &context.AnalyzerContext{
		Root:   ".",
		Dirs:   Dirs{},
		DirIdx: DirIdx{},
	}

	dirCollector = collector.DirInfoCollector{
		Ctx: ctx,
	}

	fmt.Println("collecting infos ...")
	filepath.Walk(ctx.Root, visit)

	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileHotspots)
	fmt.Printf("directory-hotspots:\n%v\n", printDirs(GetDirHotspots(ctx, topCount), printSize))
	fmt.Printf("tree-hotspots:\n%v\n\n", printDirs(GetTreeHotspots(ctx, topCount), printTotalSize))
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfo := FileInfo{path, f.Size()}
		Add(fileHotspots, fileInfo)
		dirCollector.AddFile(fileInfo)
	} else {
		dir := &Dir{Name: path, Children: []string{}}
		dirCollector.AddDir(dir)
	}
	return nil
}

func printDirs(infos Dirs, printer DirPrinter) string {
	result := ""
	for _, info := range infos {
		result += printer(info) + "\n"
	}
	return result
}

func printSize(info *Dir) string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.Size)), info.Name)
}

func printTotalSize(info *Dir) string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.TotalSize)), info.Name)
}
