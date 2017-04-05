package main

import "fmt"

func factorial(i int) int{
	if(i<=1) {
		return 1
	}
	return i * factorial(i-1)
	
}
func main(){
	i:=5;
	fmt.Println("Factorial of ",i," is:" ,factorial(i))
	
}