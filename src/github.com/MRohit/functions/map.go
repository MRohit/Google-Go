package main

import "fmt"

func main(){
	var myMap map[string]string
	myMap = make(map[string]string)
	
	myMap["India"] = "Delhi"
	myMap["Japan"] = "Tokyo"
	myMap["England"] = "London"
	myMap["Switzerland"] = "Zurich"
	
	for key := range myMap {
		fmt.Println("Key:",key," Value:",myMap[key])

	}
	
	capital, ok := myMap["United States"]
	
	if(ok) {
		fmt.Println("\nCapital:",capital)
	} else {
		fmt.Println("\nKey doesn't exist")
	}
	
	capital2, ok2 := myMap["India"]
	
	if(ok2) {
		fmt.Println("\nCapital:",capital2)
	} else {
		fmt.Println("\nKey doesn't exist")
	}
}