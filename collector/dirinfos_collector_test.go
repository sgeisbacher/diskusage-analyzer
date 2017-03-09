package collector

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func createDirInfoCollectorWithSampleCtx() DirInfoCollector {
	ctx := &AnalyzerContext{
		Root:   ".",
		Dirs:   Dirs{},
		DirIdx: DirIdx{},
	}

	col := DirInfoCollector{
		Ctx: ctx,
	}

	col.AddDir(&Dir{Name: "."})
	col.AddDir(&Dir{Name: "stefan"})
	col.AddFile(FileInfo{"stefan/file1.txt", 1000})
	col.AddFile(FileInfo{"stefan/file2.txt", 1020})
	col.AddDir(&Dir{Name: "stefan/code"})
	col.AddFile(FileInfo{"stefan/code/file1.txt", 1040})
	col.AddFile(FileInfo{"stefan/file3.txt", 1060})
	col.AddFile(FileInfo{"stefan/file4.txt", 1080})
	col.AddDir(&Dir{Name: "stefan/music"})
	col.AddFile(FileInfo{"stefan/music/song1.mp3", 1100})
	col.AddFile(FileInfo{"stefan/code/file2.txt", 1120})

	return col
}

func TestAddFileDirSizeCalc(t *testing.T) {
	RegisterTestingT(t)

	col := createDirInfoCollectorWithSampleCtx()
	ctx := col.Ctx

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

	col := createDirInfoCollectorWithSampleCtx()
	ctx := col.Ctx

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
