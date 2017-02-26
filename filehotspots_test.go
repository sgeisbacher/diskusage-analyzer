package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestFileInfosAdd(t *testing.T) {
	RegisterTestingT(t)

	fileInfos := make(FileInfos, 4)
	fileInfos.Add(FileInfo{"/file1.txt", 1000})
	fileInfos.Add(FileInfo{"/file2.txt", 1001})
	fileInfos.Add(FileInfo{"/file3.txt", 1002})
	fileInfos.Add(FileInfo{"/file5.txt", 1004})
	fileInfos.Add(FileInfo{"/file4.txt", 1003})

	Expect(len(fileInfos)).To(Equal(4))
	Expect(fileInfos[0].Name).To(Equal("/file5.txt"))
	Expect(fileInfos[1].Name).To(Equal("/file4.txt"))
	Expect(fileInfos[2].Name).To(Equal("/file3.txt"))
	Expect(fileInfos[3].Name).To(Equal("/file2.txt"))
}
