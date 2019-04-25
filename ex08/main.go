package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive only
	cs := make(chan<- int) // send only

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	go func() {
		c <- 42
		//cr <- 43 // send to receive-only type <-chan int
		cs <- 43
	}()

	fmt.Println(<-c)
	//fmt.Println(<-cr)
	//fmt.Println(<-cs) // invalid operation: <-cs (receive from send-only type chan<- int)

}
