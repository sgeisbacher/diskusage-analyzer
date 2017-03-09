package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func buildDirectoryTree() *AnalyzerContext {
	ctx := &AnalyzerContext{
		root:   "/home",
		dirs:   Dirs{},
		dirIdx: DirIdx{},
	}

	addDir(ctx, "/home", 0, []string{"/home/stefan"})
	addDir(ctx, "/home/stefan", 10, []string{"/home/stefan/code", "/home/stefan/private", "/home/stefan/multimedia"})
	addDir(ctx, "/home/stefan/code", 0, []string{"/home/stefan/code/go", "/home/stefan/code/js"})
	addDir(ctx, "/home/stefan/code/go", 250, nil)
	addDir(ctx, "/home/stefan/code/js", 350, nil)
	addDir(ctx, "/home/stefan/private", 210, []string{"/home/stefan/private/docs", "/home/stefan/private/movies"})
	addDir(ctx, "/home/stefan/private/movies", 37000, nil)
	addDir(ctx, "/home/stefan/private/docs", 20, nil)
	addDir(ctx, "/home/stefan/multimedia", 0, []string{"/home/stefan/multimedia/private"})
	addDir(ctx, "/home/stefan/multimedia/private", 0, []string{"/home/stefan/multimedia/private/photos", "/home/stefan/multimedia/private/music"})
	addDir(ctx, "/home/stefan/multimedia/private/photos", 6000, nil)
	addDir(ctx, "/home/stefan/multimedia/private/music", 0, []string{"/home/stefan/multimedia/private/music/rock", "/home/stefan/multimedia/private/music/rap", "/home/stefan/multimedia/private/music/hiphop"})
	addDir(ctx, "/home/stefan/multimedia/private/music/rock", 4000, nil)
	addDir(ctx, "/home/stefan/multimedia/private/music/rap", 500, nil)
	addDir(ctx, "/home/stefan/multimedia/private/music/hiphop", 7000, nil)

	return ctx
}

func TestCalcTotalSizes(t *testing.T) {
	RegisterTestingT(t)

	ctx := buildDirectoryTree()
	ctx.CalcTotalSizes()

	tableTestData := []struct {
		folderPath        string
		expectedTotalSize int64
	}{
		{"/home", 55340},
		{"/home/stefan", 55340},
		{"/home/stefan/code", 600},
		{"/home/stefan/code/js", 350},
		{"/home/stefan/private", 37230},
		{"/home/stefan/multimedia", 17500},
		{"/home/stefan/multimedia/private", 17500},
		{"/home/stefan/multimedia/private/music", 11500},
	}

	for _, testData := range tableTestData {
		Expect(ctx.dirIdx[testData.folderPath].TotalSize).To(Equal(testData.expectedTotalSize))
	}
}

func TestIsPotentialTreeHotspot(t *testing.T) {
	RegisterTestingT(t)

	ctx := buildDirectoryTree()
	ctx.CalcTotalSizes()

	filter := isPotentialTreeHotspot(ctx, 0.8)

	tableTestData := []struct {
		dirPath        string
		expectedResult bool
	}{
		{"/home/stefan/private", false},
		{"/home", false},
		{"/home/stefan/code", true},
		{"/home/stefan/multimedia", false},
		{"/home/stefan/multimedia/private", true},
		{"/home/stefan/private/movies", false},
	}

	for _, testData := range tableTestData {
		dir := ctx.dirIdx[testData.dirPath]
		Expect(filter(dir)).To(Equal(testData.expectedResult), fmt.Sprintf("%v should be %v", testData.dirPath, testData.expectedResult))
	}
}

func TestGetTreeHotspots(t *testing.T) {
	RegisterTestingT(t)

	ctx := buildDirectoryTree()
	hotspots := ctx.GetTreeHotspots(100)

	Expect(len(hotspots)).To(Equal(4))
	Expect(hotspots[0].Name).To(Equal("/home/stefan"))
	Expect(hotspots[1].Name).To(Equal("/home/stefan/multimedia/private"))
	Expect(hotspots[2].Name).To(Equal("/home/stefan/multimedia/private/music"))
	Expect(hotspots[3].Name).To(Equal("/home/stefan/code"))
}

func addDir(ctx *AnalyzerContext, name string, size int64, children []string) {
	dir := &Dir{Name: name, Size: size, Children: children}
	ctx.dirs = append(ctx.dirs, dir)
	ctx.dirIdx[name] = dir
}
