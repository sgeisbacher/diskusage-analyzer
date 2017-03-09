package main

import (
	"fmt"
	"math"
	"path/filepath"
	"sort"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func getOrCreateDir(ctx *AnalyzerContext, path string) *Dir {
	dir := ctx.DirIdx[path]
	if dir == nil {
		dir = &Dir{Name: path}
		ctx.DirIdx[path] = dir
		ctx.Dirs = append(ctx.Dirs, dir)
		parent, found := ctx.DirIdx[filepath.Dir(path)]
		if !found {
			fmt.Println("WARN: parent", filepath.Dir(path), "NOT FOUND")
			return dir
		}
		if parent.Children == nil {
			parent.Children = []string{}
		}
		parent.Children = append(parent.Children, path)
	}
	return dir
}

func GetDirHotspots(ctx *AnalyzerContext, top int) Dirs {
	sort.Sort(DirsSizeDescSorter(ctx.Dirs))
	limit := getLimit(len(ctx.Dirs), top)
	return ctx.Dirs[:limit]
}

func GetTreeHotspots(ctx *AnalyzerContext, top int) Dirs {
	ctx.CalcTotalSizes()
	hotspots := ctx.Dirs.Filter(isPotentialTreeHotspot(ctx, 0.8))

	sort.Sort(DirsTotalSizeDescSorter(hotspots))
	limit := getLimit(len(hotspots), top)
	return hotspots[:limit]
}

func isPotentialTreeHotspot(ctx *AnalyzerContext, threshold float64) DirFilter {
	return func(dir *Dir) bool {
		maxRelDiff := float64(0)
		if len(dir.Children) == 0 {
			return false
		}
		for _, childName := range dir.Children {
			child, found := ctx.DirIdx[childName]
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
