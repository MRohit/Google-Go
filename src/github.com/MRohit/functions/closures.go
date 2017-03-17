package main

import "fmt"

func getSequenceOfNumbers () func () int {
	i:=0
	
	return func() int {
		i+=1
		return i
	}
}

func TwoSquareRoot () func () int {
	i:=1
	
	return func () int {
		i*=2
		return i
	}
}

func TwoCubeRoot () func () int {
	i:=1
	
	return func () int {
		i*=3
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
	
	two:=TwoSquareRoot()
	for i:=1; i < 10; i++ {
		fmt.Printf("\n 2's Root :%d",two())
	}
	two:=TwoCubeRoot()
	for i:=1; i < 10; i++ {
		fmt.Printf("\n 2's Cube Root :%d",two())
	}
	
}