/*
sending a value to a channel
ch <- v

receiving a value from a channel
val := <- ch

close channel
close(ch)

query buffer of a channel{returns an integer denoting the buffer of the specified channel}
cap(ch)

length of a channel
len(ch)
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}

// sends data to the channel
func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Sent all data")
	close(ch)
	wg.Done()
}

// receives data from the chanel
func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	for val := range ch {
		fmt.Println("Recieved: ", val)
	}
	wg.Done()
}
