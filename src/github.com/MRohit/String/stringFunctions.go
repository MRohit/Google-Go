package main

import (
	"fmt"
	"strings" 
)
func main(){
	var str = "This moment"
	
	// Finding length of string using len(str)
	fmt.Printf("String:'%s' and it's length:%d\n",str,len(str))
	
	// Concatenate strings
	var str2 = []string {"This moment", "We Own it"}
	
	fmt.Printf(strings.Join(str2," "))
}