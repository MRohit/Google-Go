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
	for i:=0;i<10;i++ {
		stdArr[i].rollNo = i
		stdArr[i].name = "Student " + strconv.Itoa(i+1)
		stdArr[i].marks = (70.06 + float64(i))
	}
	
	bubbleSort(stdArr,10)
	fmt.Printf("\nSorted Student details:")
	fmt.Printf("\nRoll No\tName\tMarks\n")
	for i:=0;i<10;i++ {
		printStudeDetails(stdArr[i])
	}
}
func bubbleSort(stdArr[] student,n int) {
	swapped := 1
	for i:=0;i<n && swapped == 1;i++ {
		swapped = 0
		for j:= i+1; j < n;j++ {
			if(stdArr[i].marks < stdArr[j].marks) {
				std := stdArr[i]
				stdArr[i] = stdArr[j]
				stdArr[j] = std;
				swapped = 1
			}
		}
	}
	
}
func printStudeDetails(std1 student) {
	fmt.Printf("\n%d\t%s\t%f",std1.rollNo,std1.name,std1.marks)
}