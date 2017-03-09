package main

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func TestFileInfosAdd(t *testing.T) {
	RegisterTestingT(t)

	fileInfos := make(FileInfos, 4)
	Add(fileInfos, FileInfo{"/file1.txt", 1000})
	Add(fileInfos, FileInfo{"/file2.txt", 1001})
	Add(fileInfos, FileInfo{"/file3.txt", 1002})
	Add(fileInfos, FileInfo{"/file5.txt", 1004})
	Add(fileInfos, FileInfo{"/file4.txt", 1003})

	Expect(len(fileInfos)).To(Equal(4))
	Expect(fileInfos[0].Name).To(Equal("/file5.txt"))
	Expect(fileInfos[1].Name).To(Equal("/file4.txt"))
	Expect(fileInfos[2].Name).To(Equal("/file3.txt"))
	Expect(fileInfos[3].Name).To(Equal("/file2.txt"))
}
