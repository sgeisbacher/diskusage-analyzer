package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func createSampleCtx() *AnalyzerContext {
	ctx := &AnalyzerContext{
		root:       ".",
		dirInfos:   Dirs{},
		dirInfoIdx: DirIdx{},
	}

	ctx.AddDir(&Dir{Name: "."})
	ctx.AddDir(&Dir{Name: "stefan"})
	ctx.AddFile(FileInfo{"stefan/file1.txt", 1000})
	ctx.AddFile(FileInfo{"stefan/file2.txt", 1020})
	ctx.AddDir(&Dir{Name: "stefan/code"})
	ctx.AddFile(FileInfo{"stefan/code/file1.txt", 1040})
	ctx.AddFile(FileInfo{"stefan/file3.txt", 1060})
	ctx.AddFile(FileInfo{"stefan/file4.txt", 1080})
	ctx.AddDir(&Dir{Name: "stefan/music"})
	ctx.AddFile(FileInfo{"stefan/music/song1.mp3", 1100})
	ctx.AddFile(FileInfo{"stefan/code/file2.txt", 1120})
	return ctx
}

func TestAddFileDirSizeCalc(t *testing.T) {
	RegisterTestingT(t)

	ctx := createSampleCtx()

	Expect(len(ctx.dirInfos)).To(Equal(4))

	Expect(ctx.dirInfos[0].Name).To(Equal("."))
	Expect(ctx.dirInfos[0].Size).To(Equal(int64(0)))

	Expect(ctx.dirInfos[1].Name).To(Equal("stefan"))
	Expect(ctx.dirInfos[1].Size).To(Equal(int64(4160)))

	Expect(ctx.dirInfos[2].Name).To(Equal("stefan/code"))
	Expect(ctx.dirInfos[2].Size).To(Equal(int64(2160)))

	Expect(ctx.dirInfos[3].Name).To(Equal("stefan/music"))
	Expect(ctx.dirInfos[3].Size).To(Equal(int64(1100)))
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

	for _, dir := range ctx.dirInfos {
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
		dirInfos: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	dirHotspots := ctx.GetDirHotspots(3)

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
		dirInfos: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	for _, testData := range tableTestData {
		hotspots := ctx.GetDirHotspots(testData.topNum)
		Expect(len(hotspots)).To(Equal(testData.expectedLen))
	}
}
