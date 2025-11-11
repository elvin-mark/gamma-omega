package fft

import (
	"math/cmplx"
	"testing"
)

func isSimilar(a, b complex128) bool {
	return cmplx.Abs(a-b) <= 1.0e-6
}

func isSimilarArray(a, b []complex128) bool {
	for i := range a {
		if !isSimilar(a[i], b[i]) {
			return false
		}
	}
	return true
}
func TestFft(t *testing.T) {
	x := []complex128{1, 2, 1, 2}
	res := Fft(x)
	actualRes := []complex128{6, 0, -2, 0}
	if !isSimilarArray(res, actualRes) {
		t.Error("wrong result")
	}
}
