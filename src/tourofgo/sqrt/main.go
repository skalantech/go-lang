package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	i := 1
	for {
		z = z - (z*z-x)/(2*z)
		e := z*z - x
		fmt.Println("Iteration ", i, "e: ", e, "sqrt: ", z)
		if e < 0.0000001 {
			break
		}
		i++
	}
	return z
}

func main() {
	value := float64(30)
	fmt.Println(Sqrt(value))
	fmt.Println(math.Sqrt(value))
}
