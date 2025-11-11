package integrate

import (
	"math"
	"testing"
)

func f2(x float64) float64 {
	return x * x * x
}

func TestTrapezoid(t *testing.T) {
	s := Trapezoid(f2, 0.0, 1.0, 10)
	println(s)
	if math.Abs(s-0.25) > 0.01 {
		t.Error("did not converge")
	}
}
