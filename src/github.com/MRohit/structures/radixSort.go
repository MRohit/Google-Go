package main

import (
		"fmt" 
		)


const RANGE int = 10

func main () {
	
	var stdArr[] int
	stdArr = make([]int, 10)
	
	for i:=9;i>=0;i-- {
		stdArr[i] = i
		
	}
	
	radixSort(stdArr,10)
	
}
func getMax(stdArr[]int, n int) int {
	max := stdArr[0]
	for i:=1;i<n;i++ {
		if (max < stdArr[i]) {
			max = stdArr[i]
		}
	}
	return max
}
func radixSort(stdArr[] int,n int) {
	m := getMax(stdArr,n);
	for exp:= 1; (m/exp) > 0; exp=exp*10 {
		countingSort(stdArr,n,exp)
	}
}
func countingSort (stdArr[] int, n int,exp int) {

	var count[RANGE+1] int 
	var output[] int;
	output = make([]int,10)
	output = make([]int,n)
	for i:=0;i<=RANGE;i++ {
		count[i]=0
	}
	for i:=0;i<n;i++ {
		count[(stdArr[i]/exp)%10]++
	}
	for i:=1;i<=RANGE;i++ {
		count[i]= count[i] + count[i-1]
	}
	
	for i:=n-1;i>=0;i-- {
		output[count[(stdArr[i]/exp)%10 -1]] = stdArr[i]
		count[(stdArr[i]/exp)%10]--;
	}
	for i:=0;i<10;i++ {
		fmt.Printf("%d ",output[i])
	}
}
