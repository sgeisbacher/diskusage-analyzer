package collector

import (
	"fmt"
	"path/filepath"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type DirInfoCollector struct {
	Ctx *AnalyzerContext
}

func (col DirInfoCollector) AddFile(fileInfo FileInfo) {
	fileParent := filepath.Dir(fileInfo.Name)
	if fileParent == "." {
		return
	}
	dir, found := col.Ctx.DirIdx[fileParent]
	if !found {
		fmt.Println("WARN: dir", fileParent, "NOT FOUND")
		return
	}
	dir.Size += fileInfo.Size
}

func (col DirInfoCollector) AddDir(dir *Dir) {
	col.Ctx.Dirs = append(col.Ctx.Dirs, dir)
	col.Ctx.DirIdx[dir.Name] = dir
	if dir.Name == col.Ctx.Root {
		return
	}
	parent, found := col.Ctx.DirIdx[filepath.Dir(dir.Name)]
	if !found {
		fmt.Println("WARN: PARENT NOT FOUND:", filepath.Dir(dir.Name))
		return
	}
	parent.Children = append(parent.Children, dir.Name)
}
