// given a number x, we want to find the number z for which z^2 is most nearly x

package main

func Sqrt(x float64) float64 {
	z := 1.0
	for x > z {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
