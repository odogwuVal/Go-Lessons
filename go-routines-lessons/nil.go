package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)
	wg.Wait()
}

func sender(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Sent", i)
	}
	close(ch)
	wg.Done()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println("recieved", val)
	}
	wg.Done()
}
