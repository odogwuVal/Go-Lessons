package main

import (
	"fmt"
	"time"
)

func main() {
	func() {
		fmt.Println("Executing in anonymous mode")
	}()
	time.Sleep(1*time.Second)
	fmt.Println("Executing main function")
}