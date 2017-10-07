package main
import "fmt"
var a int = 20;
func main() {
  var a int = 100
  var b int = 10
  c := sum(a,b);
  fmt.Printf("\nValue of c in main:%d\n",c);
}

func sum(a,b int) int {
  fmt.Printf("\nValue of a in sum = %d",a)
  fmt.Printf("\nValue of b in sum=%d",b)
  return a+b;
}
