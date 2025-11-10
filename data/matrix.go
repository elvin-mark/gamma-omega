package data

import "github.com/elvin-mark/gamma-omega/utils"

type Matrix struct {
	Data []float64
	Rows int
	Cols int
}

func NewMatrix(rows int, cols int) Matrix {
	return Matrix{Data: make([]float64, rows*cols), Rows: rows, Cols: cols}
}

func (m Matrix) Zeros() {
	for i := range m.Rows * m.Cols {
		m.Data[i] = 0.0
	}
}

func (m Matrix) Ones() {
	for i := range m.Rows * m.Cols {
		m.Data[i] = 1.0
	}
}

func (m Matrix) Eye() {
	for i := range m.Rows {
		for j := range m.Cols {
			m.Data[i*m.Cols+j] = 0.0
		}
		if i < m.Cols {
			m.Data[i*m.Cols+i] = 1.0
		}
	}
}

func (m Matrix) Dot(v Vector) Vector {
	utils.Assert(m.Cols == v.Dim, "matrix cols should be the same as the dimension of the vector")
	res := NewVector(m.Rows)
	for i := range m.Rows {
		s := 0.0
		for j := range m.Cols {
			s += m.Data[i*m.Cols+j] * v.Data[j]
		}
		res.Data[i] = s
	}
	return res
}

func (m Matrix) Gemm(t Matrix) Matrix {
	utils.Assert(m.Cols == t.Rows, "matrix rows of a should be the same of b cols")
	res := NewMatrix(m.Rows, t.Cols)
	for i := range m.Rows {
		for j := range t.Cols {
			s := 0.0
			for k := range m.Cols {
				s += m.Data[i*m.Cols+k] * t.Data[k*t.Cols+j]
			}
			res.Data[i*res.Cols+j] = s
		}
	}
	return res
}

func (m Matrix) LU() (L, U Matrix) {
	utils.Assert(m.Rows == m.Cols, "matrix has to be squared")
	L = NewMatrix(m.Rows, m.Cols)
	U = NewMatrix(m.Rows, m.Cols)
	n := m.Rows
	L.Eye()
	U.Zeros()
	for i := range n {
		for j := i; j < n; j++ {
			s := 0.0
			for k := range i {
				s += L.Data[i*L.Cols+k] * U.Data[k*U.Cols+j]
			}
			U.Data[i*U.Cols+j] = m.Data[i*m.Cols+j] - s
		}
		for j := i + 1; j < n; j++ {
			s := 0.0
			for k := range i {
				s += L.Data[j*L.Cols+k] * U.Data[k*U.Cols+i]
			}
			L.Data[j*L.Cols+i] = (m.Data[j*m.Cols+i] - s) / U.Data[i*U.Cols+i]
		}
	}
	return
}

func (m Matrix) Inv() Matrix {
	utils.Assert(m.Rows == m.Cols, "matrix has to be squared")
	L, U := m.LU()
	mInv := NewMatrix(m.Rows, m.Cols)
	mInv.Zeros()
	n := m.Rows
	z := NewVector(n)
	z.Zeros()
	for i := range n {
		for j := range n {
			s := 0.0
			for k := range j {
				s += L.Data[j*L.Cols+k] * z.Data[k]
			}
			if i == j {
				z.Data[j] = 1.0 - s
			} else {
				z.Data[j] = -s
			}
		}
		for j := n - 1; j >= 0; j-- {
			s := 0.0
			for k := j + 1; k < n; k++ {
				s += U.Data[j*U.Cols+k] * mInv.Data[k*mInv.Cols+i]
			}
			mInv.Data[j*mInv.Cols+i] = (z.Data[j] - s) / U.Data[j*U.Cols+j]
		}
	}
	return mInv
}

func (m Matrix) Det() float64 {
	_, U := m.LU()
	d := 1.0
	for i := range m.Rows {
		d *= U.Data[i*U.Cols+i]
	}
	return d
}
