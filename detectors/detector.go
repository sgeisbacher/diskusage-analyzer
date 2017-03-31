package detectors

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type Hotspot struct {
	Name string
	Size int64
}

type Hotspots []Hotspot

type HotspotsDetector interface {
	GetName() string
	Detect(ctx AnalyzerContext, top int) (Hotspots, error)
}

func getLimit(size int, top int) int {
	switch {
	case top <= 0:
		return size
	case top > 0 && top <= size:
		return top
	case top > size:
		return size
	}
	return 3
}

func (hotspot Hotspot) String() string {
	return fmt.Sprintf("%10s  %v", humanize.Bytes(uint64(hotspot.Size)), hotspot.Name)
}

func (hotspots Hotspots) String() string {
	lines := []string{}
	for _, h := range hotspots {
		lines = append(lines, h.String())
	}
	return strings.Join(lines, "\n")
}
