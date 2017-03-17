package main

import "fmt"

func getSequenceOfNumbers () func () int {
	i:=0
	
	return func() int {
		i+=1
		return i
	}
}

func main(){
	nextValue :=getSequenceOfNumbers()
	
	fmt.Printf("\nNext Value:%d",nextValue())
	fmt.Printf("\nNext Value:%d",nextValue())
	fmt.Printf("\nNext Value:%d",nextValue())
	
	secondNumber :=getSequenceOfNumbers()
	fmt.Printf("\nsecondNumber Value:%d",secondNumber())
	fmt.Printf("\nsecondNumber Value:%d",secondNumber())
	fmt.Printf("\nsecondNumber Value:%d",secondNumber())
	fmt.Printf("\nsecondNumber Value:%d",secondNumber())
	
	
}