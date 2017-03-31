package main

import (
		"fmt" 
		)


const RANGE int = 255

func main () {
	
	var stdArr[] int
	stdArr = make([]int, 10)
	
	for i:=9;i>=0;i-- {
		stdArr[i] = i
		
	}
	
	countingSort(stdArr,10)
	
}
func countingSort (stdArr[] int, n int) {

	var count[RANGE+1] int 
	var output[] int;
	output = make([]int,n)
	for i:=0;i<=RANGE;i++ {
		count[i]=0
	}
	for i:=0;i<n;i++ {
		count[stdArr[i]]++
	}
	for i:=1;i<=RANGE;i++ {
		count[i]= count[i] + count[i-1]
	}
	
	for i:=n-1;i>=0;i-- {
		output[count[stdArr[i]] -1] = stdArr[i]
		count[stdArr[i]]--;
	}
	for i:=0;i<10;i++ {
		fmt.Printf("%d ",output[i])
	}
}
