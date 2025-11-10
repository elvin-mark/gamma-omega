package optimize

import (
	"fmt"
	"math"

	data "github.com/elvin-mark/gamma-omega/data"
	"github.com/elvin-mark/gamma-omega/utils"
)

func NewtonMethod(f func(float64) float64, df func(float64) float64, x0 float64) float64 {
	x := x0
	i := 0
	for math.Abs(f(x)) > utils.EPSILON && i < utils.MAX_ITERATIONS {
		x = x - f(x)/df(x)
		i += 1
	}
	if math.Abs(f(x)) <= utils.EPSILON {
		fmt.Printf("found the root after %d iterations\n", i)
	}
	return x
}

func GeneralizedNewtonMethod(f func(data.Vector) data.Vector, J func(data.Vector) data.Matrix, x0 data.Vector) data.Vector {
	i := 0
	x := x0
	for f(x).Norm() > utils.EPSILON && i < utils.MAX_ITERATIONS {
		delta := J(x).Inv().Dot(f(x))
		x = x.Sub(delta)
		i += 1
	}
	if f(x).Norm() <= utils.EPSILON {
		fmt.Printf("found the root after %d iterations\n", i)
	}
	return x
}
