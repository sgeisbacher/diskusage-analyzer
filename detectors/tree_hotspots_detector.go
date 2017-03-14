package detectors

import (
	"fmt"
	"math"
	"sort"

	humanize "github.com/dustin/go-humanize"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type TreeHotspotsDetector struct{}

func (detector TreeHotspotsDetector) GetName() string { return "tree-node hotspots" }

func (detector TreeHotspotsDetector) Detect(ctx AnalyzerContext, top int) (Dirs, error) {
	ctx.CalcTotalSizes()
	hotspots := ctx.Dirs.Filter(isPotentialTreeHotspot(ctx, 0.8))

	sort.Sort(DirsTotalSizeDescSorter(hotspots))
	limit := getLimit(len(hotspots), top)
	return hotspots[:limit], nil
}

func isPotentialTreeHotspot(ctx AnalyzerContext, threshold float64) DirFilter {
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

func (detector TreeHotspotsDetector) GetDirPrinter() DirPrinter {
	return func(info *Dir) string {
		return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(info.TotalSize)), info.Name)
	}
}
