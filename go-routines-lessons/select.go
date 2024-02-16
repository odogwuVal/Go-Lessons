package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	// wg.Add(1)
	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val2 := <-ch2:
		fmt.Println(val2)
	case val1 := <-ch1:
		fmt.Println(val1)
	default:
		fmt.Println("Exit blocks")
	}
	// wg.Wait()
	time.Sleep(1 * time.Second)
}

func goOne(ch1 chan string) {
	ch1 <- "Channel-1"
	// wg.Done()
}

func goTwo(ch2 chan string) {
	ch2 <- "Channel-2"
	// wg.Done()
}
