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
	
	fmt.Printf("Next Value:%d",nextValue())
	fmt.Printf("Next Value:%d",nextValue())
	fmt.Printf("Next Value:%d",nextValue())
	
	secondNumber :=getSequenceOfNumbers()
	fmt.Printf("secondNumber Value:%d",secondNumber())
	fmt.Printf("secondNumber Value:%d",secondNumber())
	fmt.Printf("secondNumber Value:%d",secondNumber())
	fmt.Printf("secondNumber Value:%d",secondNumber())
	
	
}