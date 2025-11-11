package linalg

import (
	"testing"

	"github.com/elvin-mark/gamma-omega/data"
)

func TestSolve(t *testing.T) {
	A := data.Matrix{
		Data: []float64{1, 2, 3, 1, 1, 1, 1, 8, 9},
		Rows: 3,
		Cols: 3,
	}
	b := data.Vector{
		Data: []float64{6, 3, 18},
		Dim:  3,
	}
	x := Solve(A, b)

	s := data.Vector{
		Data: []float64{1.0, 1.0, 1.0},
		Dim:  3,
	}
	if x.Sub(s).Norm() > 1e-6 {
		t.Error("solution is wrong")
	}

}
