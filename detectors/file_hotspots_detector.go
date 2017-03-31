package detectors

import . "github.com/sgeisbacher/diskusage-analyzer/context"

type FileHotspotsDetector struct{}

func (detector FileHotspotsDetector) GetName() string { return "file hotspots" }

func (detector FileHotspotsDetector) Detect(ctx AnalyzerContext, top int) (Hotspots, error) {
	hotspots := make(Hotspots, top)
	for i, fileInfo := range ctx.FileHotspots {
		hotspots[i] = Hotspot{fileInfo.Name, fileInfo.Size}
	}
	return hotspots, nil
}
