package data

import (
	"math"
	"testing"
)

func isSimilar(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func isMatrixSimilar(a, b Matrix) bool {
	if a.Rows != b.Rows || a.Cols != b.Cols {
		return false
	}
	for i := range a.Data {
		if !isSimilar(a.Data[i], b.Data[i]) {
			return false
		}
	}
	return true
}

func TestMatrixDet(t *testing.T) {
	A := Matrix{
		Data: []float64{
			1, 2, 3,
			4, 1, 6,
			7, 8, 9,
		},
		Rows: 3,
		Cols: 3,
	}
	d := A.Det()
	if !isSimilar(math.Abs(d-48.0), 1e-6) {
		t.Error("wrong determinant")
	}
}

func TestMatrixGemm(t *testing.T) {
	A := Matrix{
		Data: []float64{
			1, 2, 3,
			4, 5, 6,
		},
		Rows: 2,
		Cols: 3,
	}

	B := Matrix{
		Data: []float64{
			1, 2,
			3, 4,
			5, 6,
		},
		Rows: 3,
		Cols: 2,
	}
	C := A.Gemm(B)

	actualC := Matrix{
		Data: []float64{
			22, 28,
			49, 64,
		},
		Rows: 2,
		Cols: 2,
	}

	if !isMatrixSimilar(C, actualC) {
		t.Error("wrong result")
	}

}

func TestMatrixLU(t *testing.T) {
	A := Matrix{
		Data: []float64{
			1, 2, 3,
			4, 1, 6,
			7, 8, 9,
		},
		Rows: 3,
		Cols: 3,
	}
	L, U := A.LU()

	actualL := Matrix{
		Data: []float64{
			1, 0, 0,
			4, 1, 0,
			7, 0.8571428571428571, 1,
		},
		Rows: 3,
		Cols: 3,
	}

	actualU := Matrix{
		Data: []float64{
			1, 2, 3,
			0, -7, -6,
			0, 0, -6.857142857142858,
		},
		Rows: 3,
		Cols: 3,
	}

	if !isMatrixSimilar(L, actualL) {
		t.Error("wrong L")
	}

	if !isMatrixSimilar(U, actualU) {
		t.Error("wrong U")
	}

}

func TestMatrixInv(t *testing.T) {
	A := Matrix{
		Data: []float64{
			1, 2,
			3, 4,
		},
		Rows: 2,
		Cols: 2,
	}
	AInv := A.Inv()
	actualAInv := Matrix{
		Data: []float64{
			-2, 1,
			1.5, -0.5,
		},
		Rows: 2,
		Cols: 2,
	}
	if !isMatrixSimilar(AInv, actualAInv) {
		t.Error("wrong inverse")
	}
}
