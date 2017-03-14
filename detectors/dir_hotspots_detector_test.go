package detectors

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func TestGetHotspotsSorting(t *testing.T) {
	RegisterTestingT(t)

	ctx := AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	dirHotspots, err := DirHotspotsDetector{}.Detect(ctx, 3)

	Expect(err).To(BeNil())
	Expect(len(dirHotspots)).To(Equal(3))
	Expect(dirHotspots[0].Name).To(Equal("/stefan"))
	Expect(dirHotspots[1].Name).To(Equal("/stefan/code"))
	Expect(dirHotspots[2].Name).To(Equal("/stefan/music"))
}

func TestDirHotspotsDetectorDetectTopLimit(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		topNum      int
		expectedErr bool
		expectedLen int
	}{
		{topNum: 0, expectedErr: false, expectedLen: 3},
		{topNum: 2, expectedErr: false, expectedLen: 2},
		{topNum: 4, expectedErr: false, expectedLen: 3},
	}

	ctx := AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	for _, testData := range tableTestData {
		hotspots, err := DirHotspotsDetector{}.Detect(ctx, testData.topNum)
		Expect(err != nil).To(Equal(testData.expectedErr))
		Expect(len(hotspots)).To(Equal(testData.expectedLen))
	}
}
