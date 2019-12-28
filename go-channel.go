package main

import "fmt"

// We use Go Channels to connect concurrent goroutines, in order to send and receive values between them, using the channel operator.

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal

	v1, v2 := <-fooVal, <-fooVal

	fmt.Println(v1, v2)
}
