package integrate

import "github.com/elvin-mark/gamma-omega/data"

func RungeKutta(dydt func(float64, float64) float64, y0 float64, t []float64) []float64 {
	y := make([]float64, len(t))
	y[0] = y0
	for i := 1; i < len(t); i++ {
		h := t[i] - t[i-1]
		k1 := h * dydt(t[i-1], y[i-1])
		k2 := h * dydt(t[i-1]+h/2, y[i-1]+h*k1/2)
		k3 := h * dydt(t[i-1]+h/2, y[i-1]+h*k2/2)
		k4 := h * dydt(t[i-1]+h, y[i-1]+h*k3)
		y[i] = y[i-1] + (k1+2*k2+2*k3+k4)/6
	}
	return y
}

func GeneralizedRungeKutta(dydt func(float64, data.Vector) data.Vector, y0 data.Vector, t []float64) []data.Vector {
	y := make([]data.Vector, len(t))
	y[0] = y0
	for i := 1; i < len(t); i++ {
		h := t[i] - t[i-1]

		k1 := dydt(t[i-1], y[i-1]).Scale(h)
		k2 := dydt(t[i-1]+h/2, y[i-1].Add(k1.Scale(h/2.0))).Scale(h)
		k3 := dydt(t[i-1]+h/2, y[i-1].Add(k2.Scale(h/2.0))).Scale(h)
		k4 := dydt(t[i-1]+h, y[i-1].Add(k3.Scale(h))).Scale(h)

		tmp := k1.Add(k2.Scale(2)).Add(k3.Scale(2)).Add(k4).Scale(1 / 6.0)
		y[i] = y[i-1].Add(tmp)
	}
	return y
}
