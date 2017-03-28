package main

import "fmt"

func main(){
	arr:=[]int {1,2,3,4,5,6,7}
	fmt.Printf("\ncontents of Array ==%v",arr)
	
	// sub slice of array between 1 to 4 index
	fmt.Printf("\nArray[1:4]:%v",arr[1:4])
	
	subArr2 := arr[1:4]
	printSlice(subArr2)
	
	
	subArr3 := arr[4:6]
	printSlice(subArr3)
}

func printSlice(a[] int){
	fmt.Printf("\nlen=%d\tcap=%d\tslice=%v",len(a),cap(a),a)
}