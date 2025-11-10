package data

import (
	"math"

	"github.com/elvin-mark/gamma-omega/utils"
)

type Vector struct {
	Data []float64
	Dim  int
}

func NewVector(n int) Vector {
	return Vector{Data: make([]float64, n), Dim: n}
}

func (v Vector) Clone() Vector {
	res := NewVector(v.Dim)
	for i := range v.Dim {
		res.Data[i] = v.Data[i]
	}
	return res
}

func (v Vector) Zeros() {
	for i := range v.Dim {
		v.Data[i] = 0.0
	}
}

func (v Vector) Scale(s float64) {
	for i := range v.Dim {
		v.Data[i] *= s
	}
}

func (v Vector) Add(u Vector) Vector {
	utils.Assert(v.Dim == u.Dim, "vectors has to be of the same dimension")
	res := NewVector(v.Dim)
	for i := range v.Dim {
		res.Data[i] = v.Data[i] + u.Data[i]
	}
	return res
}

func (v Vector) AddInplace(u Vector) {
	utils.Assert(v.Dim == u.Dim, "vectors has to be of the same dimension")
	for i := range v.Dim {
		v.Data[i] += u.Data[i]
	}
}

func (v Vector) Sub(u Vector) Vector {
	utils.Assert(v.Dim == u.Dim, "vectors has to be of the same dimension")
	res := NewVector(v.Dim)
	for i := range v.Dim {
		res.Data[i] = v.Data[i] - u.Data[i]
	}
	return res
}

func (v Vector) Norm() float64 {
	s := 0.0
	for i := range v.Dim {
		s += v.Data[i] * v.Data[i]
	}
	return math.Sqrt(s)
}

func (v Vector) Distance(u Vector) float64 {
	return v.Sub(u).Norm()
}
