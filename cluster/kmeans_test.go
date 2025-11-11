package cluster

import (
	"testing"

	"github.com/elvin-mark/gamma-omega/data"
)

func TestKMeans(t *testing.T) {
	points := []data.Vector{
		{Data: []float64{1.0, 1.0}, Dim: 2},
		{Data: []float64{2.0, 2.0}, Dim: 2},
		{Data: []float64{3.0, 3.0}, Dim: 2},
		{Data: []float64{4.0, 4.0}, Dim: 2},
		{Data: []float64{5.0, 5.0}, Dim: 2},
		{Data: []float64{6.0, 6.0}, Dim: 2},
		{Data: []float64{7.0, 7.0}, Dim: 2},
		{Data: []float64{8.0, 8.0}, Dim: 2},
		{Data: []float64{9.0, 9.0}, Dim: 2},
		{Data: []float64{10.0, 10.0}, Dim: 2},
	}
	u, clusters := KMeans(points, 3)
	for i := range points {
		println(clusters[i])
	}
	for i := range u {
		println(u[i].Data[0], u[i].Data[1])
	}
	expectedClusters := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2}
	for i := range clusters {
		if clusters[i] != expectedClusters[i] {
			t.Error("did not converge")
		}
	}
}
