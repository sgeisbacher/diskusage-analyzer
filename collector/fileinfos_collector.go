package collector

import (
	"sort"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type FileInfoCollector struct {
	Ctx *AnalyzerContext
}

func (col FileInfoCollector) AddFile(elem FileInfo) {
	lastElemIndex := len(col.Ctx.FileHotspots) - 1
	if elem.Size > col.Ctx.FileHotspots[lastElemIndex].Size {
		col.Ctx.FileHotspots[lastElemIndex] = elem
		sort.Sort(col.Ctx.FileHotspots)
	}
}
