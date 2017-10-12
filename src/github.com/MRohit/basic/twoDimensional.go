package main
import "fmt"

func main(){
	var a = [5][2]int{ {23,19}, {1,2}, {2,4}, {3,6},{4,8}}
	var b = [5][2]int{ {2,3}, {4,2}, {3,9}, {11,12},{15,17}}
	fmt.Printf("\nMatrix A:\n")
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			fmt.Printf("%d ",a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nMatrix B:\n")
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			fmt.Printf("%d ",b[i][j])
		}
		fmt.Printf("\n")
	}
	var c[5][2] int;
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			c[i][j] = a[i][j] * b[i][j];
		}

	}


	fmt.Printf("\nMatrix multiplication result:\n")
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			fmt.Printf("%d ",c[i][j])
		}
		fmt.Printf("\n")
	}

	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j];
		}

	}


	fmt.Printf("\nMatrix Addition result:\n")
	for i:=0; i < 2; i++ {
		fmt.Printf("\n")
		for j:=0; j < 2; j++ {
			fmt.Printf("%d ",c[i][j])
		}
		fmt.Printf("\n")
	}
}
