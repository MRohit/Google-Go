package main

import "fmt"

func swap(s1,s2 string) (string,string) {
	return s2,s1
}

func main () {
	a,b:=swap("Rahul","Rohit")
	fmt.Printf("After swapping:%s and %s",a,b)
}