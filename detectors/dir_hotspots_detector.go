package detectors

import (
	"errors"
	"sort"

	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type DirHotspotsDetector struct{}

func (detector DirHotspotsDetector) GetName() string { return "directory hotspots" }

func (detector DirHotspotsDetector) Detect(ctx AnalyzerContext, top int) (Hotspots, error) {
	if len(ctx.Dirs) == 0 {
		return nil, errors.New("no data to analyze")
	}
	sort.Sort(DirsSizeDescSorter(ctx.Dirs))
	limit := getLimit(len(ctx.Dirs), top)
	hotspots := make(Hotspots, limit)
	for i := 0; i < limit; i++ {
		dir := ctx.Dirs[i]
		hotspots[i] = Hotspot{Name: dir.Name, Size: dir.Size}
	}
	return hotspots, nil
}
