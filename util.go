package np

import (
	"errors"
	"math"
)

func root(a float64, n int) float64 {
	n1 := n - 1
	n1f, rn := float64(n1), 1/float64(n)
	x, x0 := 1., 0.
	for {
		potx, t2 := 1/x, a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 *= potx
			}
			potx *= potx
		}
		x0, x = x, rn*(n1f*x+t2)
		if math.Abs(x-x0)*1e15 < x {
			break
		}
	}
	return x
}

func validateArray(err error, size int) error {
	if err != nil {
		return err
	}

	if size == 0 {
		return errors.New(ErrEmptyArray)
	}

	return nil
}

func float64Pointer(value float64) *float64 {
	return &value
}

func intPointer(value int) *int {
	return &value
}

func int32Pointer(value int32) *int32 {
	return &value
}
