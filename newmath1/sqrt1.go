// name of a package rather than running it as main
package newmath1

// sqrt returns an approxmiation fo the squareroot of x.
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z * z / x) / (2 * z)
	}
	return z
}
