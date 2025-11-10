package integrate

func Trapezoid(f func(float64) float64, a float64, b float64, n int) float64 {
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	h := (b - a) / float64(n)
	for i := range x {
		x[i] = a + float64(i)*h
		y[i] = f(x[i])
	}
	s := 0.0
	for i := 1; i < n; i++ {
		s += y[i]
	}
	s *= 2
	s += y[0] + y[n]
	s *= h / 2
	return s
}
