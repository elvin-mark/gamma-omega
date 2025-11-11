package fft

import (
	"math"

	"github.com/elvin-mark/gamma-omega/utils"
)

func isPowerofTwo(n int) bool {
	return n&(n-1) == 0
}

func Fft(x []complex128) []complex128 {
	N := len(x)
	if N == 1 {
		return x
	}
	utils.Assert(isPowerofTwo(N), "length of data has to be a power of two")
	even := make([]complex128, N/2)
	odd := make([]complex128, N/2)

	for i := range x {
		if i%2 == 0 {
			even[i/2] = x[i]
		} else {
			odd[i/2] = x[i]
		}
	}

	evenFft := Fft(even)
	oddFft := Fft(odd)

	res := make([]complex128, N)

	for k := range N / 2 {
		theta := 2 * math.Pi * float64(k) / float64(N)
		t := complex(math.Cos(theta), -math.Sin(theta))
		res[k] = evenFft[k] + oddFft[k]*t
	}

	for k := range N / 2 {
		theta := 2 * math.Pi * float64(k) / float64(N)
		t := complex(math.Cos(theta), -math.Sin(theta))
		res[k+N/2] = evenFft[k] - oddFft[k]*t
	}

	return res

}
