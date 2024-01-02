package difference_check

import (
	"testing"
)

func TestPercentChange(t *testing.T) {
	result := PercentChange(1, 2)

	if result != 100 {
		t.Errorf("Wrong result. Expect 100. Actual %f", result)
	}
}

func TestPercentDifference(t *testing.T) {
	result := PercentDifference(1, 3)

	if result != float64(100) {
		t.Errorf("Wrong result. Expect 100. Actual %f", result)
	}
}
