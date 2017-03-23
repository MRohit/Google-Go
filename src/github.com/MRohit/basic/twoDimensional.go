package main
import "fmt"

func main(){
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	fmt.Printf("\nTwo Dimensional Array:\n")
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			fmt.Printf("%d ",a[i][j])
		}
		fmt.Printf("\n")
	}
}