package main

import (
	"fmt"
	)

type MyPrint interface {
	print()
}

type result struct {
	status int
}

type error struct {
	msg string
}

func (res result) print(){
	fmt.Println("Status:",res.status)
}
func (err error) print(){
	fmt.Println(err.msg)
}
func main(){
	 res:= result{status:200}
	 err := error{msg:"Error occurred"}
	
	res.print()
	err.print()
}