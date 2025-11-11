package special

import (
	"testing"
)

func TestLegendreP(t *testing.T) {
	x := 1.5
	res := LegendreP(4, x)
	if !isSimilar(res, 14.0859375) {
		t.Error("wrong result")
	}
}
