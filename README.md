# go-difference-check

This package checks how relatively large one value is to another. The functions are used for estimation, so inputs and results are in float64 type. Result is rounded to number of decimal places based on the given precision by user.
It can be used for price comparison.

Run 
```
  go get github.com/sydneyshile/go-difference-check
```
to download package. 

Functions:

1. PercentChange(initial, final float64, precision uint) (float64, error)
2. PercentDifference(n1, n2 float64, precision uint) (float64, error)