package data

import "testing"

func TestVectorNorm(t *testing.T) {
	v := NewVector(2)
	v.Data[0] = 3.0
	v.Data[1] = 4.0
	n := v.Norm()
	if n != 5.0 {
		t.Error("wrong norm")
	}
}
