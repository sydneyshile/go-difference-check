package difference

import (
    "errors"
    "math"
)

/*
PercentChange This function takes change from initial to final value, dividing it by absolute value of initial value,
and multiplying it by 100.
If result is positive, there is an increase.
If result is negative, there is a decrease.
Initial value cannot be zero.
*/
func PercentChange(initial float64, final float64) (float64, error) {
    if initial == 0 {
        return 0, errors.New("Initial value should be non-zero")
    }
    diff := final - initial
    return diff / math.Abs(initial) * 100, nil
}

/*
PercentDifference This function takes two numbers and choses average value as reference point.
It calculates absolute difference, divides it with average value, and multiplies it
by 100 to get percent difference.
Both values must be greater than zero.
*/
func PercentDifference(n1 float64, n2 float64) (float64, error) {
    if n1 <= 0 || n2 <= 0 {
        return 0, errors.New("Values should be greater than 0")
    }
    diff := math.Abs(n1 - n2)
    average := (n1 + n2) / 2
    return (diff / average * 100), nil
}
