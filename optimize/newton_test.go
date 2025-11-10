package optimize

import (
	"fmt"
	"math"
	"testing"

	data "github.com/elvin-mark/gamma-omega/data"
	"github.com/elvin-mark/gamma-omega/utils"
)

func TestNewton(t *testing.T) {
	f := func(x float64) float64 {
		return x * math.Cos(x)
	}

	df := func(x float64) float64 {
		return math.Cos(x) - x*math.Sin(x)
	}

	x := NewtonMethod(f, df, 2.0)
	fmt.Println(x)
	fmt.Println(f(x))
	if math.Abs(f(x)) > utils.EPSILON {
		panic("did not converge")
	}
}

func TestGeneralizedNewton(t *testing.T) {
	f2d := func(v data.Vector) data.Vector {
		res := data.NewVector(2)
		x := v.Data[0]
		y := v.Data[1]
		res.Data[0] = x*x + y*y - 4
		res.Data[1] = x - y
		return res
	}

	J2d := func(v data.Vector) data.Matrix {
		res := data.NewMatrix(2, 2)
		res.Zeros()
		res.Data[0] = 2 * v.Data[0]
		res.Data[1] = 2 * v.Data[1]
		res.Data[2] = 1
		res.Data[3] = -1
		return res
	}

	x := GeneralizedNewtonMethod(f2d, J2d, data.Vector{Data: []float64{1.0, 1.0}, Dim: 2})
	fmt.Println(x)
	fmt.Println(f2d(x))
	if f2d(x).Norm() > utils.EPSILON {
		panic("did not converge")
	}
}
