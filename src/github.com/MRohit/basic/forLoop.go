package main

import "fmt"

func main(){
	var x int = 15
	var y int
	// defining an array
	numbers :=[6] int {1,2,3,4,5,6}
	
	// sample for loop
	for y:=0; y < 10;y++  {
		fmt.Printf("Value of y:%d\n",y)
	}
	
	// Conditional Loop
	for x<y{
		x++
		fmt.Printf("Value of x:%d\n",x)
	}
	// Range Loop, kind for foreach loop
	for i,z:= range numbers {
		fmt.Printf("Value of z=:%d at %d\n",z,i)
	}
}