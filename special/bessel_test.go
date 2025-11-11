package special

import (
	"fmt"
	"testing"
)

func TestBesselJ0(t *testing.T) {
	x := 1.0
	res := BesselJ0(x)
	if !isSimilar(res, 0.7651976865579665) {
		t.Error("wrong result")
	}
}

func TestBesselJ1(t *testing.T) {
	x := 1.0
	res := BesselJ1(x)
	if !isSimilar(res, 0.44005058574493355) {
		t.Error("wrong result")
	}
}

func TestBesselJ4(t *testing.T) {
	x := 1.0
	res := BesselJ(4, x)
	fmt.Println(res)
	if !isSimilar(res, 0.002476638964109955) {
		t.Error("wrong result")
	}
}
