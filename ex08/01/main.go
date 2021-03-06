package main

import "fmt"

func main() {

	c := make(chan int)

	//send channel
	go foo(c)

	// receive channel
	bar(c)

	println("about to exit")
}

//send
func foo(c chan<- int) { // general to specific (send only)
	c <- 42
	close(c) //close the channel
}

//receive
func bar(c <-chan int) { // general to specific (receive only)
	v, ok := <-c // comma ok idiom
	println(v, ok)
	fmt.Printf("type is %T\n", c)
	println(<-c) //pull off after close
}
