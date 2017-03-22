package main

import "fmt"

func main(){
	var str = "First String"
	fmt.Printf("Value of str:%s\n",str)
	
	fmt.Printf("\nHex Bytes:")
	for i:=0; i < len(str); i++ {
		fmt.Printf("%x ",str[i])
	}
	fmt.Printf("\n")
   
   const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
   /*q flag escapes unprintable characters, with + flag it escapses non-ascii characters as well 
     to make output unambigous  */
   fmt.Printf("quoted string: ")
   fmt.Printf("%+q", sampleText)
   
   fmt.Printf("\nquoted string: ")
   fmt.Printf("%+flag", sampleText)
   fmt.Printf("\n")  
}