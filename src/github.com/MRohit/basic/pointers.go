package main

import "fmt"

const MAX int = 3

func main() {
	var a  = [] int { 10, 20, 30 }
	var ptr[MAX]*int;
	for i:=0;i<MAX;i++ {
		ptr[i] = &a[i];
	}
	
	for i:=0;i<MAX;i++ {
		fmt.Printf("\nValue of  a[%d] = %d",i,*ptr[i])
	}
	
	var b int = 5
	var bptr* int;
	var	bptr2** int;
	bptr = &b
	bptr2 = &bptr 
	
	fmt.Printf("\nb Value:%d",b)
	fmt.Printf("\nSingle Pointer bptr:%d",*bptr)
	fmt.Printf("\nDouble Pointer bptr2:%d",**bptr2)
}