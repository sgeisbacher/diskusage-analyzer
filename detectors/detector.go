package detectors

import (
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

type HotspotsDetector interface {
	GetName() string
	Detect(ctx AnalyzerContext, top int) (Dirs, error)
	GetDirPrinter() DirPrinter
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
