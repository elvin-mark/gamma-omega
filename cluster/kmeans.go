package cluster

import (
	data "github.com/elvin-mark/gamma-omega/data"
	utils "github.com/elvin-mark/gamma-omega/utils"
)

func KMeans(points []data.Vector, n int) (u []data.Vector, clusters []int) {
	utils.Assert(len(points) <= n, "number of clusters cannot be greater than the number of points")
	u = make([]data.Vector, n)
	clusters = make([]int, len(points))
	for i := range n {
		u[i] = points[i].Clone()
	}
	for i := 0; i < utils.MAX_ITERATIONS; i++ {
		for j := range points {
			clusters[j] = 0
			d := u[0].Distance(points[j])
			for k := 1; k < n; k++ {
				tmp := u[k].Distance(points[j])
				if tmp < d {
					clusters[j] = k
					d = tmp
				}
			}
		}
		for k := 0; k < n; k++ {
			u[k].Zeros()
			count := 0
			for j := range points {
				if clusters[j] == k {
					u[k].AddInplace(points[j])
					count++
				}
			}
			u[k].Scale(1.0 / float64(count))
		}
	}
	return u, clusters
}
