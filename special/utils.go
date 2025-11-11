package special

import "math"

func isSimilar(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}
