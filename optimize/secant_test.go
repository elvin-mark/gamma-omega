package optimize

import (
	"fmt"
	"math"
	"testing"

	"github.com/elvin-mark/gamma-omega/utils"
)

func TestSecant(t *testing.T) {
	f := func(x float64) float64 {
		return x * math.Cos(x)
	}

	x := SecantMethod(f, 1.0, 2.0)
	fmt.Println(x)
	fmt.Println(f(x))
	if math.Abs(f(x)) > utils.EPSILON {
		t.Error("did not converge")
	}
}
