package difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentChange(t *testing.T) {
	type test struct {
		testname      string
		initial       float64
		final         float64
		expected      float64
		expectederror string
	}

	tests := []test{
		{testname: "Both integers", initial: 1, final: 2, expected: 100, expectederror: ""},
		{testname: "Both decimals", initial: 100.0, final: 150.0, expected: 50, expectederror: ""},
		{testname: "One negative inputs", initial: -1, final: 12, expected: 1300, expectederror: ""},
		{testname: "Two negative inputs", initial: -1, final: -10, expected: -900, expectederror: ""},
		{testname: "Zero for initial value", initial: 0, final: 5, expected: 0, expectederror: "Initial value should be non-zero"},
	}

	for _, tc := range tests {
		actual, actualerr := PercentChange(tc.initial, tc.final)
		assert.Equal(t, tc.expected, actual, "Test name: "+tc.testname)
		if tc.expectederror != "" {
			assert.EqualError(t, actualerr, tc.expectederror)
		}
	}
}

func TestPercentDifference(t *testing.T) {
	type test struct {
		testname      string
		n1            float64
		n2            float64
		expected      float64
		expectederror string
	}

	tests := []test{
		{testname: "Both integers", n1: 1, n2: 2.5, expected: 85.71428571428571, expectederror: ""},
		{testname: "Both decimals", n1: 100.0, n2: 150.0, expected: 40, expectederror: ""},
		{testname: "Negative numbers", n1: -1, n2: 12, expected: 0, expectederror: "Values should be greater than 0"},
		{testname: "Zero inputs", n1: -1, n2: -10, expected: 0, expectederror: "Values should be greater than 0"},
	}

	for _, tc := range tests {
		actual, actualerr := PercentDifference(tc.n1, tc.n2)
		assert.Equal(t, tc.expected, actual, "Test name: "+tc.testname)
		if tc.expectederror != "" {
			assert.EqualError(t, actualerr, tc.expectederror)
		}
	}
}
