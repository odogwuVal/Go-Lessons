package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sendValue(ch)

	select {
	case msg := <-ch:
		fmt.Println("Received", msg)
	case <-time.After(1 * time.Second):
		fmt.Println(("select timeout"))
	}
}

func sendValue(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 20
}
