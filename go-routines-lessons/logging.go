package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/valentinemadu/GO-ROUTINES/logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(file)
	log.Println("Hello World!")
}
