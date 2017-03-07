package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func buildDirectoryTree() *DirHotspotsContext {
	ctx := &DirHotspotsContext{
		root:       "/home",
		dirInfos:   DirInfos{},
		dirInfoIdx: DirInfoIdx{},
	}

	addDirInfo(ctx, "/home", 0, []string{"/home/stefan"})
	addDirInfo(ctx, "/home/stefan", 10, []string{"/home/stefan/code", "/home/stefan/private", "/home/stefan/multimedia"})
	addDirInfo(ctx, "/home/stefan/code", 0, []string{"/home/stefan/code/go", "/home/stefan/code/js"})
	addDirInfo(ctx, "/home/stefan/code/go", 250, nil)
	addDirInfo(ctx, "/home/stefan/code/js", 350, nil)
	addDirInfo(ctx, "/home/stefan/private", 210, []string{"/home/stefan/private/docs", "/home/stefan/private/movies"})
	addDirInfo(ctx, "/home/stefan/private/movies", 37000, nil)
	addDirInfo(ctx, "/home/stefan/private/docs", 20, nil)
	addDirInfo(ctx, "/home/stefan/multimedia", 0, []string{"/home/stefan/multimedia/private"})
	addDirInfo(ctx, "/home/stefan/multimedia/private", 0, []string{"/home/stefan/multimedia/private/photos", "/home/stefan/multimedia/private/music"})
	addDirInfo(ctx, "/home/stefan/multimedia/private/photos", 6000, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music", 0, []string{"/home/stefan/multimedia/private/music/rock", "/home/stefan/multimedia/private/music/rap", "/home/stefan/multimedia/private/music/hiphop"})
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/rock", 4000, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/rap", 500, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/hiphop", 7000, nil)

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
		Expect(ctx.dirInfoIdx[testData.folderPath].TotalSize).To(Equal(testData.expectedTotalSize))
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
		dirInfo := ctx.dirInfoIdx[testData.dirPath]
		Expect(filter(dirInfo)).To(Equal(testData.expectedResult), fmt.Sprintf("%v should be %v", testData.dirPath, testData.expectedResult))
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

func addDirInfo(ctx *DirHotspotsContext, name string, size int64, children []string) {
	dirInfo := &DirInfo{Name: name, Size: size, Children: children}
	ctx.dirInfos = append(ctx.dirInfos, dirInfo)
	ctx.dirInfoIdx[name] = dirInfo
}
