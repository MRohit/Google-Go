package main

import (
		"fmt" 
		"strconv"
		)

type student struct {
	rollNo int
	name string
	marks float64
}

func main () {
	
	var stdArr[] student
	stdArr = make([]student, 10)
	var temp[] student
	temp = make([]student, 10)
	for i:=0;i<10;i++ {
		stdArr[i].rollNo = i
		stdArr[i].name = "Student " + strconv.Itoa(i+1)
		stdArr[i].marks = (70.06 + float64(i))
	}
	
	mergeSort(stdArr,temp,0,9)
	fmt.Printf("\nSorted Student details:")
	fmt.Printf("\nRoll No\tName\tMarks\n")
	for i:=0;i<10;i++ {
		printStudeDetails(stdArr[i])
	}
}
func mergeSort(stdArr[] student,temp[] student,left int,right int) {
	var mid int
	if(right > left) {
		mid = (right+left)/2	//l+(r-l)/2;
		mergeSort(stdArr,temp,left,mid)
		mergeSort(stdArr,temp,mid+1,right)
		merge(stdArr,temp,left,mid,right)
	}	
}
func merge(stdArr[] student, temp[]student, left int, mid int, right int) {
	size:=(right-left)+1
	tempPos:=left
	leftEnd:=mid-1
	
	for left <= leftEnd && mid <= right {
		if(stdArr[left].marks > stdArr[mid].marks) {
			temp[tempPos] = stdArr[left]
			left++
			
		} else {
			temp[tempPos] = stdArr[mid]
			mid++
			
		}
		tempPos++
	}
	
	for left <= leftEnd {
		temp[tempPos] = stdArr[left]
		left++
		tempPos++
	}
	
	for mid <= right {
		temp[tempPos] = stdArr[mid]
		mid++
		tempPos++
	}
	
	for i:=0; i<size;i++ {
		
		stdArr[right] = temp[right]
		right--
	}
}
func printStudeDetails(std1 student) {
	fmt.Printf("\n%d\t%s\t%f",std1.rollNo,std1.name,std1.marks)
}