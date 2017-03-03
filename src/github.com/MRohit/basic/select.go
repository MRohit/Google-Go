package main

import "fmt"
import "time"
func main(){
	c1 :=make(chan int)
	c2 :=make(chan int)
	c3 :=make(chan int)
	var i1,i2 int
	i1=1
	go func() {
        time.Sleep(time.Second * 1)
        c1 <- 1
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- 2
    }()
	go func() {
        time.Sleep(time.Second * 3)
        c3 <- 3
    }()
	
	for i := 0; i < 3; i++ {
		select {
			case i1 = <-c1:
				fmt.Printf("recieved ",i1, "from c1\n")
			case c2 <-i2:
				fmt.Printf("sent ", i2, " to c2\n")
			case i3, ok := (<-c3):
				if ok {
					fmt.Printf("received ", i3, " from c3\n")
				 } else {
					fmt.Printf("c3 is closed\n")
				 }
			default:
				fmt.Printf("no communication\n")
		}
	}
	
	c11 := make(chan string)
    c21 := make(chan string)
	//Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
    go func() {
        time.Sleep(time.Second * 1)
        c11 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c21 <- "two"
    }()
	
	//We’ll use select to await both of these values simultaneously, printing each one as it arrives.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c11:
            fmt.Println("received", msg1)
        case msg2 := <-c21:
            fmt.Println("received", msg2)
        }
    }
}