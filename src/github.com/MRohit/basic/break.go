package main
import "fmt"

func main(){
	var a int = 5
	for a < 25 {
		fmt.Printf("Value of a:%d\n",a)
		a++
		if(a==20) {
			break;
		}
	}
}