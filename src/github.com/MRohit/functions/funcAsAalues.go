package main

import (
	"fmt"
	"math"
)
func main(){
	getSquareRoot := func(a float64) float64 {
		return math.Sqrt(a)
	}
	getAddition := func(a,b float64) float64 {
		return a+b
	}
	
	fmt.Printf("\nSquare root of 169 is %f",getSquareRoot(169))
	
	fmt.Printf("\nAddition of 10 & 20 is %f",getAddition(10,20))
}