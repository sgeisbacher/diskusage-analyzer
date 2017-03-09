package main

import "fmt"
import "sort"
import "github.com/dustin/go-humanize"

type FileInfo struct {
	Name string
	Size int64
}

type FileInfos []FileInfo

func (infos FileInfos) Len() int           { return len(infos) }
func (infos FileInfos) Swap(i, j int)      { infos[i], infos[j] = infos[j], infos[i] }
func (infos FileInfos) Less(i, j int) bool { return infos[i].Size > infos[j].Size }

func (infos FileInfos) Add(elem FileInfo) {
	size := len(infos)
	if elem.Size > infos[size-1].Size {
		infos[size-1] = elem
		sort.Sort(infos)
	}
}

func (infos FileInfos) String() string {
	result := ""
	for _, info := range infos {
		result += info.String() + "\n"
	}
	return result
}

func (info FileInfo) String() string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.Size)), info.Name)
}
