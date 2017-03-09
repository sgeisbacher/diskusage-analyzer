package main

type Dirs []*Dir
type DirIdx map[string]*Dir

type AnalyzerContext struct {
	root   string
	dirs   Dirs
	dirIdx DirIdx
}

func (ctx *AnalyzerContext) CalcTotalSizes() {
	ctx.calcTotalSizes(ctx.root)
}

func (ctx *AnalyzerContext) calcTotalSizes(path string) int64 {
	dir, found := ctx.dirIdx[path]
	if !found {
		return 0
	}
	sum := dir.Size
	for _, childPath := range dir.Children {
		sum += ctx.calcTotalSizes(childPath)
	}
	dir.TotalSize = sum
	return sum
}
