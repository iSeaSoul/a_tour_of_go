package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	loop_tm := 0
	for true {
		new_z := z - (z * z - x) / (2 * z)
		if math.Abs(new_z - z) < 1e-16 {
			return z
		}
		z = new_z
		loop_tm += 1
		if loop_tm > 100 {
			return z
		}
	}
	return -1
}

func main() {
	fmt.Println(Sqrt(5))
}
