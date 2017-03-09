package main

type Dirs []*Dir
type DirIdx map[string]*Dir

type AnalyzerContext struct {
	root       string
	dirInfos   Dirs
	dirInfoIdx DirIdx
}

func (ctx *AnalyzerContext) CalcTotalSizes() {
	ctx.calcTotalSizes(ctx.root)
}

func (ctx *AnalyzerContext) calcTotalSizes(path string) int64 {
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
