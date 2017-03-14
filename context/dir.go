package context

type Dir struct {
	Name      string
	TotalSize int64
	Size      int64
	Children  []string
}

type DirPrinter func(info *Dir) string

type DirsSizeDescSorter []*Dir
type DirsTotalSizeDescSorter []*Dir

func (dir DirsSizeDescSorter) Len() int           { return len(dir) }
func (dir DirsSizeDescSorter) Swap(i, j int)      { dir[i], dir[j] = dir[j], dir[i] }
func (dir DirsSizeDescSorter) Less(i, j int) bool { return dir[i].Size > dir[j].Size }

func (dir DirsTotalSizeDescSorter) Len() int           { return len(dir) }
func (dir DirsTotalSizeDescSorter) Swap(i, j int)      { dir[i], dir[j] = dir[j], dir[i] }
func (dir DirsTotalSizeDescSorter) Less(i, j int) bool { return dir[i].TotalSize > dir[j].TotalSize }

type DirFilter func(dir *Dir) bool

func (vs Dirs) Filter(f DirFilter) Dirs {
	vsf := make(Dirs, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
