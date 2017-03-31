package detectors

import (
	"fmt"
	"math"
	"sort"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type TreeHotspotsDetector struct{}

func (detector TreeHotspotsDetector) GetName() string { return "tree-node hotspots" }

func (detector TreeHotspotsDetector) Detect(ctx AnalyzerContext, top int) (Hotspots, error) {
	ctx.CalcTotalSizes()
	topTreeNodes := ctx.Dirs.Filter(isPotentialTreeHotspot(ctx, 0.8))

	sort.Sort(DirsTotalSizeDescSorter(topTreeNodes))
	limit := getLimit(len(topTreeNodes), top)

	hotspots := make(Hotspots, limit)
	for i := 0; i < limit; i++ {
		treeNode := topTreeNodes[i]
		hotspots[i] = Hotspot{Name: treeNode.Name, Size: treeNode.TotalSize}
	}
	return hotspots, nil
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
