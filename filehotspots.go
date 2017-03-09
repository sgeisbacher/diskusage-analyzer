package main

import (
	"sort"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func Add(infos FileInfos, elem FileInfo) {
	size := len(infos)
	if elem.Size > infos[size-1].Size {
		infos[size-1] = elem
		sort.Sort(infos)
	}
}
