package main

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sgeisbacher/diskusage-analyzer/context"
)

func TestGetHotspotsSorting(t *testing.T) {
	RegisterTestingT(t)

	ctx := &AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	dirHotspots := GetDirHotspots(ctx, 3)

	Expect(len(dirHotspots)).To(Equal(3))
	Expect(dirHotspots[0].Name).To(Equal("/stefan"))
	Expect(dirHotspots[1].Name).To(Equal("/stefan/code"))
	Expect(dirHotspots[2].Name).To(Equal("/stefan/music"))
}

func TestGetLimit(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		size     int
		top      int
		expected int
	}{
		{size: 5, top: -1, expected: 5},
		{size: 5, top: 0, expected: 5},
		{size: 5, top: 3, expected: 3},
		{size: 5, top: 5, expected: 5},
		{size: 5, top: 6, expected: 5},
		{size: 5, top: 7, expected: 5},
	}

	for _, testData := range tableTestData {
		limit := getLimit(testData.size, testData.top)
		Expect(limit).To(Equal(testData.expected))
	}
}

func TestGetHotspotsTopLimit(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		topNum      int
		expectedLen int
	}{
		{topNum: 0, expectedLen: 3},
		{topNum: 2, expectedLen: 2},
		{topNum: 4, expectedLen: 3},
	}

	ctx := &AnalyzerContext{
		Dirs: Dirs{
			&Dir{"/stefan/music", 0, 1000, nil},
			&Dir{"/stefan", 0, 1100, nil},
			&Dir{"/stefan/code", 0, 1020, nil},
		},
	}

	for _, testData := range tableTestData {
		hotspots := GetDirHotspots(ctx, testData.topNum)
		Expect(len(hotspots)).To(Equal(testData.expectedLen))
	}
}
