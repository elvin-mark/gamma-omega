package optimize

import (
	"fmt"
	"math"

	"github.com/elvin-mark/gamma-omega/utils"
)

func SecantMethod(f func(float64) float64, a float64, b float64) float64 {
	i := 0
	x := b
	for math.Abs(f(x)) > utils.EPSILON && i < utils.MAX_ITERATIONS {
		x = (a*f(b) - b*f(a)) / (f(b) - f(a))
		a = b
		b = x
		i += 1
	}
	fmt.Printf("found the root after %d iterations\n", i)
	return x

}
