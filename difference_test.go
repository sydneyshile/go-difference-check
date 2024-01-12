package difference

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPercentChange(t *testing.T) {
    result := PercentChange(1, 2)

    if result != 100 {
        t.Errorf("Wrong result. Expect 100. Actual %f", result)
    }
}

func TestPercentDifference(t *testing.T) {
    result := PercentDifference(1, 3)
    expected := float64(100)
    assert.Equal(t, expected, result)
}

func TestPercentDifferenceDecimal(t *testing.T) {
    result := PercentDifference(1, 2.5)
    expected := float64(85.71428571428571)

    assert.Equal(t, expected, result)
}
