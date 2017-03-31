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
	ctx           *AnalyzerContext
	fileCollector collector.FileInfoCollector
	dirCollector  collector.DirInfoCollector
	topCount      int
)

func init() {
	flag.IntVar(&topCount, "n", 10, "limit top-hotspots count")
}

func main() {
	flag.Parse()

	ctx = &context.AnalyzerContext{
		Root:         ".",
		FileHotspots: make(FileInfos, topCount),
		Dirs:         Dirs{},
		DirIdx:       DirIdx{},
	}

	fileCollector = collector.FileInfoCollector{Ctx: ctx}
	dirCollector = collector.DirInfoCollector{Ctx: ctx}

	hotspotsDetectors := []detectors.HotspotsDetector{
		detectors.FileHotspotsDetector{},
		detectors.DirHotspotsDetector{},
		detectors.TreeHotspotsDetector{},
	}

	fmt.Println("collecting infos ...")
	filepath.Walk(ctx.Root, visit)

	fmt.Println("analyzing ...")
	for _, detector := range hotspotsDetectors {
		hotspots, err := detector.Detect(*ctx, topCount)
		if err != nil {
			fmt.Printf("%v: %v", detector.GetName(), err)
			continue
		}

		fmt.Printf("%v:\n%v\n\n", detector.GetName(), hotspots)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		fileInfo := FileInfo{path, f.Size()}
		fileCollector.AddFile(fileInfo)
		dirCollector.AddFile(fileInfo)
	} else {
		dir := &Dir{Name: path, Children: []string{}}
		dirCollector.AddDir(dir)
	}
	return nil
}
