package main

import (
	"fmt"
	"flag" // allows to parse commandline arguments
	"os"	// gives access to system calls
	"io/ioutil" // package for http
	"net/http"
	"crypto/tls"
	"io"
	"strconv"
	"strings"
	"html/net"
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

	links := All(resp.Body)  

	for _, link := range(links) {  
		fmt.Println(link)            
	}
}

func All(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}