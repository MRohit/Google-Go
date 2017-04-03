package main

import (
	"fmt"
	"flag" // allows to parse commandline arguments
	"os"	// gives access to system calls
	)
func main(){
	flag.Parse()
	args := flag.Args()
	
	if (len(args) < 1) {
		fmt.Println("Specify any URL")
		os.Exit(1)		
	}
}