package difference

import (
	"errors"
	"math"
)

/*
PercentChange This function takes change from initial to final value, dividing it by absolute value of initial value,
and multiplying it by 100.
It rounds result to specific decimal place according to the value entered for precision in argument.
If result is positive, there is an increase.
If result is negative, there is a decrease.
Initial value cannot be zero.
*/
func PercentChange(initial, final float64, precision uint) (float64, error) {
	if initial == 0 {
		return 0, errors.New("Initial value should be non-zero")
	}
	diff := final - initial
	result := diff / math.Abs(initial) * 100
	return roundFloat(result, precision), nil
}

/*
PercentDifference This function takes two numbers and choses average value as reference point.
It calculates absolute difference, divides it with average value, and multiplies it
It rounds result to specific decimal place according to the value entered for precision in argument.
Both values must be greater than zero.
*/
func PercentDifference(n1, n2 float64, precision uint) (float64, error) {
	if n1 <= 0 || n2 <= 0 {
		return 0, errors.New("Values should be greater than 0")
	}
	diff := math.Abs(n1 - n2)
	average := (n1 + n2) / 2
	result := (diff / average * 100)
	return roundFloat(result, precision), nil
}

/*
roundFloat is a helper function that rounds a value of float64 type to a specific precision.
*/
func roundFloat(value float64, precision uint) float64 {
	divisor := math.Pow(10, float64(precision))
	return math.Round(value*divisor) / divisor
}
