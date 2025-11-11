package special

import (
	"math"

	"github.com/elvin-mark/gamma-omega/utils"
)

func BesselJ0(x float64) float64 {
	sum := 0.0
	term := 1.0
	x2 := (x / 2) * (x / 2)

	for k := 0; k < utils.MAX_ITERATIONS; k++ {
		if k > 0 {
			term *= -x2 / float64(k*k)
		}
		sum += term
		if math.Abs(term) < utils.EPSILON {
			break
		}
	}
	return sum
}

func BesselJ1(x float64) float64 {
	sum := 0.0
	term := x / 2
	x2 := (x / 2) * (x / 2)

	for k := 0; k < utils.MAX_ITERATIONS; k++ {
		if k > 0 {
			term *= -x2 / float64(k*(k+1))
		}
		sum += term
		if math.Abs(term) < utils.EPSILON {
			break
		}
	}
	return sum
}

func BesselJ(n int, x float64) float64 {
	if n == 0 {
		return BesselJ0(x)
	}
	if n == 1 {
		return BesselJ1(x)
	}
	j0 := BesselJ0(x)
	j1 := BesselJ1(x)
	for k := 1; k < n; k++ {
		j2 := (2*float64(k)/x)*j1 - j0
		j0, j1 = j1, j2
	}
	return j1
}
