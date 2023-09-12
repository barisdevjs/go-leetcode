package main

import "math"

func MySqrt(x int) int {
	f := float64(x)
	return int(math.Sqrt(f))
}

func MySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func MySqrt2(x int) int {
	var out = 0

	for out*out < x {
		out++
	}

	if out*out > x {
		out--
	}

	return out
}
