package main

import "fmt"

func main(){
	var a int = 100
	var b int = 200;
	
	fmt.Printf("Greatest number is :%d",max(a,b))
}

func max (num1,num2 int) int {
	if (num1 > num2) {
		return num1
	} else	{
		return num2
	}
	
}