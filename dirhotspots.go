package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/dustin/go-humanize"
)

type DirInfo struct {
	Name      string
	TotalSize int64
	Size      int64
}

type DirInfos []*DirInfo
type DirInfoIdx map[string]*DirInfo
type DirHotspotsContext struct {
	dirInfos   DirInfos
	dirInfoIdx DirInfoIdx
}

func (infos DirInfos) Len() int           { return len(infos) }
func (infos DirInfos) Swap(i, j int)      { infos[i], infos[j] = infos[j], infos[i] }
func (infos DirInfos) Less(i, j int) bool { return infos[i].Size > infos[j].Size }

func (infos DirInfos) String() string {
	result := ""
	for _, info := range infos {
		result += info.String() + "\n"
	}
	return result
}

func (info DirInfo) String() string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.Size)), info.Name)
}

func (ctx *DirHotspotsContext) Add(fileInfo FileInfo) {
	fileParent := filepath.Dir(fileInfo.Name)
	dirInfo := ctx.getOrCreateDirInfo(fileParent)
	dirInfo.Size += fileInfo.Size
}

func (ctx *DirHotspotsContext) getOrCreateDirInfo(path string) *DirInfo {
	dirInfo := ctx.dirInfoIdx[path]
	if dirInfo == nil {
		dirInfo = &DirInfo{Name: path}
		ctx.dirInfoIdx[path] = dirInfo
		ctx.dirInfos = append(ctx.dirInfos, dirInfo)
	}
	return dirInfo
}

func (ctx *DirHotspotsContext) GetHotspots(top int) DirInfos {
	sort.Sort(ctx.dirInfos)
	limit := getLimit(len(ctx.dirInfos), top)
	return ctx.dirInfos[:limit]
}

func getLimit(size int, top int) int {
	switch {
	case top <= 0:
		return size
	case top > 0 && top <= size:
		return top
	case top > size:
		return size
	}
	return 3
}
