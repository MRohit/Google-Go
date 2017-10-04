package main
import "fmt"

func main(){
	var marks int = 80
	var grade string = "B"
	switch marks{
		case 95:
			grade = "A+"
		case 80:
			grade = "A"
		case 50,60,70:
			grade = "B"
		case 30:
			grade = "C"
	}
	switch {
		case grade == "A+":
			fmt.Printf("Excellent\n")

		case grade == "A":
			fmt.Printf("Very Good\n")

		case grade == "B":
			fmt.Printf("You Passed\n")

		case grade == "C":
			fmt.Printf("You Loser\n")
		default:
         fmt.Printf("Invalid grade\n" );


	}
	fmt.Printf("Your grade is  %s\n", grade );
}
