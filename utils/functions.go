package utils

func Linspace(a float64, b float64, n int) []float64 {
	res := make([]float64, n)
	h := (b - a) / float64(n)
	for i := range res {
		res[i] = a + float64(i)*h
	}
	return res
}
