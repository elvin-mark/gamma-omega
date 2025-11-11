package special

func LegendreP(n int, x float64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	p0 := 1.0
	p1 := x
	for k := 1; k < n; k++ {
		p2 := ((2*float64(k)+1)*x*p1 - float64(k)*p0) / (float64(k) + 1)
		p0, p1 = p1, p2
	}
	return p1
}
