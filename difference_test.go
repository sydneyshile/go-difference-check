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
		{testname: "One decimal", initial: 12.5, final: 0, expected: -100, expectederror: ""},
		{testname: "Both decimals", initial: 100.0, final: 150.0, expected: 50, expectederror: ""},
		{testname: "Both decimals 1", initial: 12.55, final: 50.456745345, expected: 302.0457796414343, expectederror: ""},
		{testname: "Both decimals 2", initial: 302.045779641, final: 50.4567453478, expected: -83.29500070890879, expectederror: ""},
		{testname: "Both decimals 3", initial: 4846545465456.888, final: 2155655.09, expected: -99.9999555218226, expectederror: ""},
		{testname: "One negative integers", initial: -1, final: 12, expected: 1300, expectederror: ""},
		{testname: "Two negative integers", initial: -1, final: -10, expected: -900, expectederror: ""},
		{testname: "On negative decimal", initial: 100.0, final: -150.0, expected: -250, expectederror: ""},
		{testname: "Both negative decimals", initial: -12.55, final: -50.456745345, expected: -302.0457796414343, expectederror: ""},
		{testname: "Both negative decimals 1", initial: -302.045, final: -50.4567453478, expected: 83.29495758982934, expectederror: ""},
		{testname: "Both negative decimals 2", initial: -48465454654.8889, final: -2155655.09, expected: 99.99555218225983, expectederror: ""},
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
		{testname: "Both integers", n1: 1, n2: 3, expected: 100, expectederror: ""},
		{testname: "One decimal", n1: 1, n2: 2.5, expected: 85.71428571428571, expectederror: ""},
		{testname: "Both decimals", n1: 100.0, n2: 150.0, expected: 40, expectederror: ""},
		{testname: "Both decimals 1", n1: 12.55, n2: 50.456745345, expected: 120.32599093140797, expectederror: ""},
		{testname: "Both decimals 2", n1: 302.045779641, n2: 50.4567453478, expected: 142.7445288802931, expectederror: ""},
		{testname: "Both decimals 3", n1: 4846545465456.888, n2: 2155655.09, expected: 199.99982208736952, expectederror: ""},
		{testname: "One negative number", n1: -1, n2: 12, expected: 0, expectederror: "Values should be greater than 0"},
		{testname: "Two negative numbers", n1: -1, n2: -10, expected: 0, expectederror: "Values should be greater than 0"},
	}

	for _, tc := range tests {
		actual, actualerr := PercentDifference(tc.n1, tc.n2)
		assert.Equal(t, tc.expected, actual, "Test name: "+tc.testname)
		if tc.expectederror != "" {
			assert.EqualError(t, actualerr, tc.expectederror)
		}
	}
}
