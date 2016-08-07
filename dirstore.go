package main

import "fmt"
import "path/filepath"
import "github.com/dustin/go-humanize"

type DirInfo struct {
	Name      string
	TotalSize int64
	Size      int64
}

type DirInfos map[string]DirInfo

func (dirInfo DirInfo) String() string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(dirInfo.Size)), dirInfo.Name)
}

func (dirInfos DirInfos) Add(fileInfo FileInfo) {
	if len(dirInfos) > 10 {
		return
	}
	fileParent := filepath.Dir(fileInfo.Name)
	dirInfo := dirInfos[fileParent]
	dirInfo.Name = fileParent
	dirInfo.Size += fileInfo.Size
	dirInfos[fileParent] = dirInfo
}
