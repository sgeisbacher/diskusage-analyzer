package main

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func createSampleCtx() *AnalyzerContext {
	ctx := &AnalyzerContext{
		Root:   ".",
		Dirs:   Dirs{},
		DirIdx: DirIdx{},
	}

	AddDir(ctx, &Dir{Name: "."})
	AddDir(ctx, &Dir{Name: "stefan"})
	AddFile(ctx, FileInfo{"stefan/file1.txt", 1000})
	AddFile(ctx, FileInfo{"stefan/file2.txt", 1020})
	AddDir(ctx, &Dir{Name: "stefan/code"})
	AddFile(ctx, FileInfo{"stefan/code/file1.txt", 1040})
	AddFile(ctx, FileInfo{"stefan/file3.txt", 1060})
	AddFile(ctx, FileInfo{"stefan/file4.txt", 1080})
	AddDir(ctx, &Dir{Name: "stefan/music"})
	AddFile(ctx, FileInfo{"stefan/music/song1.mp3", 1100})
	AddFile(ctx, FileInfo{"stefan/code/file2.txt", 1120})
	return ctx
}

func TestAddFileDirSizeCalc(t *testing.T) {
	RegisterTestingT(t)

	ctx := createSampleCtx()

	Expect(len(ctx.Dirs)).To(Equal(4))

	Expect(ctx.Dirs[0].Name).To(Equal("."))
	Expect(ctx.Dirs[0].Size).To(Equal(int64(0)))

	Expect(ctx.Dirs[1].Name).To(Equal("stefan"))
	Expect(ctx.Dirs[1].Size).To(Equal(int64(4160)))

	Expect(ctx.Dirs[2].Name).To(Equal("stefan/code"))
	Expect(ctx.Dirs[2].Size).To(Equal(int64(2160)))

	Expect(ctx.Dirs[3].Name).To(Equal("stefan/music"))
	Expect(ctx.Dirs[3].Size).To(Equal(int64(1100)))
}

func TestAddDirChildren(t *testing.T) {
	RegisterTestingT(t)

	ctx := createSampleCtx()

	expectedChildrenMap := map[string][]string{
		".":            {"stefan"},
		"stefan":       {"stefan/code", "stefan/music"},
		"stefan/code":  {},
		"stefan/music": {},
	}

	for _, dir := range ctx.Dirs {
		expectedChildren := expectedChildrenMap[dir.Name]
		Expect(len(dir.Children)).To(Equal(len(expectedChildren)), dir.Name)
		for i, expectedChild := range expectedChildren {
			Expect(dir.Children[i]).To(Equal(expectedChild))
		}
	}
}

func TestGetHotspotsSorting(t *testing.T) {
	RegisterTestingT(t)

	ctx := &AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	dirHotspots := GetDirHotspots(ctx, 3)

	Expect(len(dirHotspots)).To(Equal(3))
	Expect(dirHotspots[0].Name).To(Equal("/stefan"))
	Expect(dirHotspots[1].Name).To(Equal("/stefan/code"))
	Expect(dirHotspots[2].Name).To(Equal("/stefan/music"))
}

func TestGetLimit(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		size     int
		top      int
		expected int
	}{
		{size: 5, top: -1, expected: 5},
		{size: 5, top: 0, expected: 5},
		{size: 5, top: 3, expected: 3},
		{size: 5, top: 5, expected: 5},
		{size: 5, top: 6, expected: 5},
		{size: 5, top: 7, expected: 5},
	}

	for _, testData := range tableTestData {
		limit := getLimit(testData.size, testData.top)
		Expect(limit).To(Equal(testData.expected))
	}
}

func TestGetHotspotsTopLimit(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		topNum      int
		expectedLen int
	}{
		{topNum: 0, expectedLen: 3},
		{topNum: 2, expectedLen: 2},
		{topNum: 4, expectedLen: 3},
	}

	ctx := &AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	for _, testData := range tableTestData {
		hotspots := GetDirHotspots(ctx, testData.topNum)
		Expect(len(hotspots)).To(Equal(testData.expectedLen))
	}
}
