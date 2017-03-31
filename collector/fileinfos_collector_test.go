package collector

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func TestFileInfosAdd(t *testing.T) {
	RegisterTestingT(t)

	fileInfos := make(FileInfos, 4)
	col := FileInfoCollector{
		Ctx: &AnalyzerContext{
			FileHotspots: fileInfos,
		},
	}
	col.AddFile(FileInfo{"/file1.txt", 1000})
	col.AddFile(FileInfo{"/file2.txt", 1001})
	col.AddFile(FileInfo{"/file3.txt", 1002})
	col.AddFile(FileInfo{"/file5.txt", 1004})
	col.AddFile(FileInfo{"/file4.txt", 1003})

	Expect(len(fileInfos)).To(Equal(4))
	Expect(fileInfos[0].Name).To(Equal("/file5.txt"))
	Expect(fileInfos[1].Name).To(Equal("/file4.txt"))
	Expect(fileInfos[2].Name).To(Equal("/file3.txt"))
	Expect(fileInfos[3].Name).To(Equal("/file2.txt"))
}
