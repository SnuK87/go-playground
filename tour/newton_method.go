package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 2.0
	
	for i := 1; i < 50; i++ {
		a := z * z - x
		b := 2 * z
		c := a / b
		znew := z - c		
		delta := znew - z

		if math.Abs(delta) < 1e-6 {
			fmt.Println("skip at", i)
			return z
		}
		
		z = znew;
	}

	return z
}

func main() {

	i := 1.0
	for i < 11 {
		fmt.Println("Newton method for", i)
		a := Sqrt(i)
		e := math.Sqrt(i)
		fmt.Println("Approximated: ", a)
		fmt.Println("Expected: ", e)
		fmt.Println("Delta: ", e - a)
		i++
		fmt.Println("-----------------------")
	}
}
