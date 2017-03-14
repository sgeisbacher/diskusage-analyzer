package detectors

import (
	"errors"
	"fmt"
	"sort"

	humanize "github.com/dustin/go-humanize"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type DirHotspotsDetector struct{}

func (detector DirHotspotsDetector) GetName() string { return "directory hotspots" }

func (detector DirHotspotsDetector) Detect(ctx AnalyzerContext, top int) (Dirs, error) {
	if len(ctx.Dirs) == 0 {
		return nil, errors.New("no data to analyze")
	}
	sort.Sort(DirsSizeDescSorter(ctx.Dirs))
	limit := getLimit(len(ctx.Dirs), top)
	return ctx.Dirs[:limit], nil
}

func (detector DirHotspotsDetector) GetDirPrinter() DirPrinter {
	return func(info *Dir) string {
		return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.Size)), info.Name)
	}
}
