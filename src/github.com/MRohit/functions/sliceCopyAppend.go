package main

import "fmt"

func main(){
	var num[] int
	fmt.Printf("\nNil Array:")
	printSlice(num)
	fmt.Printf("\nAfter Appending elements one by one:")
	num = append(num,25)
	printSlice(num)
	
	num = append(num,35)
	printSlice(num)
	
	num = append(num,45)
	printSlice(num)
	num = append(num,55)
	printSlice(num)
	num = append(num,65)
	printSlice(num)
	num = append(num,75)
	printSlice(num)
	
	fmt.Printf("\nAfter copy:")
	num2 := make([]int, len(num), (cap(num))*2)
	copy(num2,num)
	
	printSlice(num2)
}

func printSlice(a[] int){
	fmt.Printf("\nlen=%d\tcap=%d\tslice=%v",len(a),cap(a),a)
}