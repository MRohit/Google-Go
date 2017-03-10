package main

import (
	"fmt"
	"math"
)
func main(){
	getSquareRoot := func(a float64) float64 {
		return math.Sqrt(a)
	}
	
	fmt.Printf("\nSquare root of 169 is %f",getSquareRoot(169))
}