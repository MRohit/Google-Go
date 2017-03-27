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
	var std1 student
	var stdArr[10] student
	std1.rollNo = 1
	std1.name = "Rohit"
	std1.marks = 70.01
	for i:=0;i<10;i++ {
		stdArr[i].rollNo = i
		stdArr[i].name = "Student " + strconv.Itoa(i+1)
		stdArr[i].marks = (70.06 + float64(i))
	}
	fmt.Printf("\nStudent details:")
	fmt.Printf("\nRoll No\tName\tMarks\n")
	for i:=0;i<10;i++ {
		printStudeDetails(stdArr[i])
	}
	
	//fmt.Printf("\n%d\t%s\t%f",std1.rollNo,std1.name,std1.marks)
}

func printStudeDetails(std1 student) {
	fmt.Printf("\n%d\t%s\t%f",std1.rollNo,std1.name,std1.marks)
}