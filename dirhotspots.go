package main

import (
	"fmt"
	"math"
	"path/filepath"
	"sort"

	"github.com/dustin/go-humanize"
)

type DirInfo struct {
	Name      string
	TotalSize int64
	Size      int64
	Children  []string
}

type DirInfos []*DirInfo
type DirInfosSizeDescSorter []*DirInfo
type DirInfosTotalSizeDescSorter []*DirInfo
type DirInfoIdx map[string]*DirInfo
type DirHotspotsContext struct {
	root       string
	dirInfos   DirInfos
	dirInfoIdx DirInfoIdx
}

func (dir DirInfosSizeDescSorter) Len() int           { return len(dir) }
func (dir DirInfosSizeDescSorter) Swap(i, j int)      { dir[i], dir[j] = dir[j], dir[i] }
func (dir DirInfosSizeDescSorter) Less(i, j int) bool { return dir[i].Size > dir[j].Size }

func (dir DirInfosTotalSizeDescSorter) Len() int           { return len(dir) }
func (dir DirInfosTotalSizeDescSorter) Swap(i, j int)      { dir[i], dir[j] = dir[j], dir[i] }
func (dir DirInfosTotalSizeDescSorter) Less(i, j int) bool { return dir[i].TotalSize > dir[j].TotalSize }

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

func (ctx *DirHotspotsContext) GetDirHotspots(top int) DirInfos {
	sort.Sort(DirInfosSizeDescSorter(ctx.dirInfos))
	limit := getLimit(len(ctx.dirInfos), top)
	return ctx.dirInfos[:limit]
}

type DirInfoFilter func(di *DirInfo) bool

func (ctx *DirHotspotsContext) GetTreeHotspots(top int) DirInfos {
	ctx.CalcTotalSizes()
	hotspots := ctx.dirInfos.Filter(isPotentialTreeHotspot(ctx, 0.8))

	sort.Sort(DirInfosTotalSizeDescSorter(hotspots))
	limit := getLimit(len(hotspots), top)
	return hotspots[:limit]
}

func isPotentialTreeHotspot(ctx *DirHotspotsContext, threshold float64) DirInfoFilter {
	return func(dir *DirInfo) bool {
		maxRelDiff := float64(0)
		if len(dir.Children) == 0 {
			return false
		}
		for _, childName := range dir.Children {
			child, found := ctx.dirInfoIdx[childName]
			if !found {
				fmt.Printf("warn: child '%v' not found in index!!!!\n", childName)
				continue
			}
			relDiff := float64(child.TotalSize) / float64(dir.TotalSize)
			maxRelDiff = math.Max(maxRelDiff, relDiff)
			if maxRelDiff > threshold {
				return false
			}
		}
		return true
	}
}

func (vs DirInfos) Filter(f DirInfoFilter) DirInfos {
	vsf := make(DirInfos, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (ctx *DirHotspotsContext) CalcTotalSizes() {
	ctx.calcTotalSizes(ctx.root)
}

func (ctx *DirHotspotsContext) calcTotalSizes(path string) int64 {
	dirInfo, found := ctx.dirInfoIdx[path]
	if !found {
		return 0
	}
	sum := dirInfo.Size
	for _, childPath := range dirInfo.Children {
		sum += ctx.calcTotalSizes(childPath)
	}
	dirInfo.TotalSize = sum
	return sum
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
