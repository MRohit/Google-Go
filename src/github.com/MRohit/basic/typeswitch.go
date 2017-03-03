package main

import "fmt"

func main(){
	var a interface{}
	
	switch i := a.(type) {
		case nil:	  
         fmt.Printf("type of x :%T",i) 
		case int:
			fmt.Printf("a is of type int\n")
		case float64:
			fmt.Printf("a is of type float64\n")
		case func(int) float64:
			fmt.Printf("a is of type func(int)\n")
		case bool,string:
			fmt.Printf("a is of type bool or string\n")
		default:
			fmt.Printf("a is of unknown type\n")
	}
}