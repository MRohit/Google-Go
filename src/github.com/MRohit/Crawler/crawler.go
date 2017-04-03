package main

import (
	"fmt"
	"flag" // allows to parse commandline arguments
	"os"	// gives access to system calls
	"io/ioutil" // package for http
	"net/http"
	"crypto/tls"
	)
func main(){
	flag.Parse()
	args := flag.Args()
	
	if (len(args) < 1) {
		fmt.Println("Specify any URL")
		os.Exit(1)		
	}
	
	getPageDetails(args[0])
}
func getPageDetails(uri string) {
	
	resp, err := http.Get(uri)
	
	if (err != nil) {
		fmt.Println("Http error thrown:",err)
		return
	}
	// Close the connection
	defer resp.Body.Close()
	
	// convert response into string
	body ,_ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}