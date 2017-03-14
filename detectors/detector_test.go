package detectors

import (
	"testing"

	. "github.com/onsi/gomega"
)

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
