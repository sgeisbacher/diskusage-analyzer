package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestTotalSizeCalculation(t *testing.T) {
	RegisterTestingT(t)

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
	addDirInfo(ctx, "/home/stefan/private", 210, []string{"/home/stefan/private/movies", "/home/stefan/private/tvshows", "/home/stefan/private/docs"})
	addDirInfo(ctx, "/home/stefan/private/movies", 25000, nil)
	addDirInfo(ctx, "/home/stefan/private/tvshows", 12000, nil)
	addDirInfo(ctx, "/home/stefan/private/docs", 20, nil)
	addDirInfo(ctx, "/home/stefan/multimedia", 0, []string{"/home/stefan/multimedia/private"})
	addDirInfo(ctx, "/home/stefan/multimedia/private", 0, []string{"/home/stefan/multimedia/private/photos", "/home/stefan/multimedia/private/music"})
	addDirInfo(ctx, "/home/stefan/multimedia/private/photos", 6000, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music", 0, []string{"/home/stefan/multimedia/private/music/rock", "/home/stefan/multimedia/private/music/rap", "/home/stefan/multimedia/private/music/hiphop"})
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/rock", 4000, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/rap", 500, nil)
	addDirInfo(ctx, "/home/stefan/multimedia/private/music/hiphop", 7000, nil)

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

func addDirInfo(ctx *DirHotspotsContext, name string, size int64, children []string) {
	dirInfo := &DirInfo{Name: name, Size: size, Children: children}
	ctx.dirInfos = append(ctx.dirInfos, dirInfo)
	ctx.dirInfoIdx[name] = dirInfo
}
