package linalg

import "github.com/elvin-mark/gamma-omega/data"

func Solve(A data.Matrix, b data.Vector) data.Vector {
	return A.Inv().Dot(b)
}
