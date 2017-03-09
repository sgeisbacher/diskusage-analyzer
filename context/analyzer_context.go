package context

type Dirs []*Dir
type DirIdx map[string]*Dir

type AnalyzerContext struct {
	Root   string
	Dirs   Dirs
	DirIdx DirIdx
}

func (ctx *AnalyzerContext) CalcTotalSizes() {
	ctx.calcTotalSizes(ctx.Root)
}

func (ctx *AnalyzerContext) calcTotalSizes(path string) int64 {
	dir, found := ctx.DirIdx[path]
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
