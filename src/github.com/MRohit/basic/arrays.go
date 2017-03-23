package main

import "fmt"

func main(){
	var a [10] int
	
	for i:=0; i < 10; i++ {
		a[i] = i + 10
	}
	fmt.Printf("\nArray Contents:")
	for i:=0; i < 10; i++ {
		fmt.Printf("%d ",a[i])
	}
	fmt.Printf("\n")
}