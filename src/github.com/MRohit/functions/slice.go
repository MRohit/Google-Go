package main

import "fmt"

func main (){
	var a[] int
	a=make([]int,3,5)
	printSlice(a)
	
	var a2[] int
	if(a2 == nil){
		fmt.Printf("\nSlice is nil")
	}
	printSlice(a2)
}
func printSlice(val []int) {
	fmt.Printf("\nlen=%d\tcap=%d\tslice=%v",len(val),cap(val),val)
}