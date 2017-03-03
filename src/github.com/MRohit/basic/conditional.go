package main
import "fmt"

func main(){
	var a int = 15
	
	if(a<15){
		fmt.Printf("a is smaller than 15\n")
	} else {
		fmt.Printf("a is greater that 15\n")
	}
	
	fmt.Printf("Value of a:%d\n",a)
}