package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sgeisbacher/diskusage-analyzer/collector"
	"github.com/sgeisbacher/diskusage-analyzer/context"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
	"github.com/sgeisbacher/diskusage-analyzer/detectors"
)

var (
	fileHotspots FileInfos
	ctx          *AnalyzerContext
	dirCollector collector.DirInfoCollector
	topCount     int
)

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

	hotspotsDetectors := []detectors.HotspotsDetector{
		detectors.DirHotspotsDetector{},
		detectors.TreeHotspotsDetector{},
	}

	fmt.Println("collecting infos ...")
	filepath.Walk(ctx.Root, visit)

	fmt.Println("analyzing ...")
	fmt.Printf("file-hotspots:\n%v\n", fileHotspots)
	for _, detector := range hotspotsDetectors {
		hotDirs, err := detector.Detect(*ctx, topCount)
		if err != nil {
			fmt.Printf("%v: %v", detector.GetName(), err)
			continue
		}
		fmt.Printf("%v:\n%v\n", detector.GetName(), printDirs(hotDirs, detector.GetDirPrinter()))
	}
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
