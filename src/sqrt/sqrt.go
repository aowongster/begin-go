package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x / 2) // the trick is to have a z thats reasonably close
	for y:= 0; y < 10000; y ++ {
		// fmt.Println(z)
		z = z - (math.Pow(z,2) -x) / 2*z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
