package difference

import (
	"math"
)

/*
This function takes absolute change between two numbers, dividing it by the initial value,
and multiplying it by 100.
*/
func PercentChange(initial float64, final float64) float64 {
	diff := math.Abs(final - initial)

	return diff / initial * 100
}

/*
This function takes two numbers and choses average value as reference point.
It calculates absolute difference and divides it by average value to get percent difference.
Minus sign infront of resulting number can be ignored as we take absolute value of difference.
*/
func PercentDifference(n1 float64, n2 float64) float64 {
	diff := math.Abs(n1 - n2)
	average := math.Abs((n1 + n2) / 2)
	return (diff/average * 100)
}
