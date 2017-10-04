package main

import "fmt"

func main(){
	var a, b,c int;
	b = 10
	a=5
	c=a+b
	fmt.Print("Addition is:")
	fmt.Println(c);
	fmt.Printf("c is of type %T\n",c);
	var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)
   y := 40.0
   fmt.Println(y)
   fmt.Printf("y is of type %T\n", y)

	 // dynamic type inference for assigning data type at runtime
	 var d,e,f = 3,3.4,"f string"
	 fmt.Println("d is of type %T",d)
	 fmt.Println("e is of type %T",e)
	 fmt.Println("f is of type %T",f)
}
