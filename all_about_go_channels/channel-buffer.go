package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
        // If you add one more element to the channel then goroutine receives the channel data
        c <- 4

	fmt.Println("main() stopped")
}

