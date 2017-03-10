package main

import "fmt"

func main(){
	var a int = 1000
	var b int = 2000
	
	fmt.Printf("\nBefore swap function a = %d\tb = %d",a,b)
	swap(&a,&b)
	
	fmt.Printf("\nAfter swap function a = %d\tb = %d",a,b)
	
}

func swap(a* int,b* int){

}