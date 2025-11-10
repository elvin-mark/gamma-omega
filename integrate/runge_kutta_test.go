package integrate

import (
	"math"
	"testing"

	data "github.com/elvin-mark/gamma-omega/data"
	"github.com/elvin-mark/gamma-omega/utils"
)

func TestRungeKutta(t *testing.T) {
	dydt := func(t float64, y float64) float64 {
		return y
	}
	tt := utils.Linspace(0.0, 1.0, 100)
	y := RungeKutta(dydt, 1.0, tt)
	for i := range y {
		println(y[i])
	}
	if math.Abs(y[99]-math.Exp(1)) > 0.1 {
		panic("did not converge")
	}
}

func TestGeneralizedRungeKutta(t *testing.T) {
	dydt := func(t float64, y data.Vector) data.Vector {
		res := data.NewVector(2)
		res.Data[0] = y.Data[1]
		res.Data[1] = -y.Data[0]
		return res
	}
	tt := utils.Linspace(0.0, 1.0, 100)
	y := GeneralizedRungeKutta(dydt, data.Vector{Data: []float64{1.0, 0.0}, Dim: 2}, tt)
	for i := range y {
		println(y[i].Data[0], y[i].Data[1])
	}
	if math.Abs(y[99].Data[0]-math.Cos(1)) > 0.1 {
		panic("did not converge")
	}
	if math.Abs(-y[99].Data[1]-math.Sin(1)) > 0.1 {
		panic("did not converge")
	}
}
