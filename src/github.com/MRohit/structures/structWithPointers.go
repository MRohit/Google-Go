package main

import (
		"fmt"
		)

type language struct {
	id int
	name string
	features string
}

func main () {
	var lg1 language
	var languagePtr *language
	lg1.id = 1
	lg1.name = "C"
	lg1.features = "Procedural, case sensitive"
	languagePtr = &lg1

	fmt.Printf("\nLanguage details:")
	fmt.Printf("\nId\tName\tFeatures\n")
	printStudeDetails(languagePtr);

}

func printStudeDetails(lg1* language) {
	fmt.Printf("\n%d\t%s\t%s",lg1.id,lg1.name,lg1.features)
}
