package utils

import "math"

func IsSimilar(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func Linspace(a float64, b float64, n int) []float64 {
	res := make([]float64, n)
	h := (b - a) / float64(n)
	for i := range res {
		res[i] = a + float64(i)*h
	}
	return res
}
