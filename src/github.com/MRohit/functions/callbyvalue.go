package main

import "fmt"

func main() {
   var a int = 10
   var b int = 20

   fmt.Printf("Before swap, value of a : %d\n", a )
   fmt.Printf("Before swap, value of b : %d\n", b )
   swap(a, b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )
}
func swap(x, y int) {
   var temp int

   temp = x 
   x = y    
   y = temp 

}