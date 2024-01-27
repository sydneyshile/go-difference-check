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
		precision     uint
		expected      float64
		expectederror string
	}

	tests := []test{
		{testname: "Both integers", initial: 1, final: 2, precision: 7, expected: 100, expectederror: ""},
		{testname: "One decimal", initial: 12.5, final: 0, precision: 1, expected: -100, expectederror: ""},
		{testname: "Both decimals", initial: 100.0, final: 150.0, precision: 3, expected: 50, expectederror: ""},
		{testname: "Both decimals 1", initial: 12.55, final: 50.456745345, precision: 5, expected: 302.04578, expectederror: ""},
		{testname: "Both decimals 2", initial: 302.045779641, final: 50.4567453478, precision: 0, expected: -83, expectederror: ""},
		{testname: "Both decimals 3", initial: 4846545465456.888, final: 2155655.09, precision: 10, expected: -99.9999555218, expectederror: ""},
		{testname: "One negative integers", initial: -1, final: 12, precision: 5, expected: 1300, expectederror: ""},
		{testname: "Two negative integers", initial: -1, final: -10, precision: 3, expected: -900, expectederror: ""},
		{testname: "On negative decimal", initial: 100.0, final: -150.0, precision: 1, expected: -250, expectederror: ""},
		{testname: "Both negative decimals", initial: -12.55, final: -50.456745345, precision: 4, expected: -302.0458, expectederror: ""},
		{testname: "Both negative decimals 1", initial: -302.045, final: -50.4567453478, precision: 0, expected: 83, expectederror: ""},
		{testname: "Both negative decimals 2", initial: -48465454654.8889, final: -2155655.09, precision: 4, expected: 99.9956, expectederror: ""},
		{testname: "Zero for initial value", initial: 0, final: 5, precision: 2, expected: 0, expectederror: "Initial value should be non-zero"},
	}

	for _, tc := range tests {
		actual, actualerr := PercentChange(tc.initial, tc.final, tc.precision)
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
		precision     uint
		expected      float64
		expectederror string
	}

	tests := []test{
		{testname: "Both integers", n1: 1, n2: 3, precision: 0, expected: 100, expectederror: ""},
		{testname: "One decimal", n1: 1, n2: 2.5, precision: 2, expected: 85.71, expectederror: ""},
		{testname: "Both decimals", n1: 100.0, n2: 150.0, precision: 2, expected: 40, expectederror: ""},
		{testname: "Both decimals 1", n1: 12.55, n2: 50.456745345, precision: 4, expected: 120.326, expectederror: ""},
		{testname: "Both decimals 2", n1: 302.045779641, n2: 50.4567453478, precision: 5, expected: 142.74453, expectederror: ""},
		{testname: "Both decimals 3", n1: 4846545465456.888, n2: 2155655.09, precision: 6, expected: 199.999822, expectederror: ""},
		{testname: "One negative number", n1: -1, n2: 12, precision: 2, expected: 0, expectederror: "Values should be greater than 0"},
		{testname: "Two negative numbers", n1: -1, n2: -10, precision: 100000, expected: 0, expectederror: "Values should be greater than 0"},
	}

	for _, tc := range tests {
		actual, actualerr := PercentDifference(tc.n1, tc.n2, tc.precision)
		assert.Equal(t, tc.expected, actual, "Test name: "+tc.testname)
		if tc.expectederror != "" {
			assert.EqualError(t, actualerr, tc.expectederror)
		}
	}
}
