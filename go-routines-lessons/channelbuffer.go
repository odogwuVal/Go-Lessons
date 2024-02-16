// a channel that needs a receiver as soon as a message is emmited to the channel is called unbuffered channel
// a buffered channel has some capacity to hold data
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, s *sync.WaitGroup) {
	ch <- 10
	ch <- 15
	ch <- 20
	fmt.Println("length: ", len(ch))
	go buy(ch, s)
	s.Done()
}

func buy(ch chan int, s *sync.WaitGroup) {
	fmt.Println("Recieved data: ", <-ch)
	fmt.Println("Recieved data: ", <-ch)
	fmt.Println("Recieved data: ", <-ch)
	s.Done()
}
