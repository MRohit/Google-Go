package main

import (
   "fmt"
   "math"
)

// Define a structure
type Circle struct {
	x,y,radius float64
}

// define method which will be associated with structure
func (circle Circle) area() float64 {
	return math.Pi*circle.radius * circle.radius
}

func main(){
	// create structure reference
	circle := Circle {x:0,y:0,radius:5}
	fmt.Printf("Circle area:%f", circle.area())
}