package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/valentinemadu/GO-ROUTINES/gofile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("\nI am testing for new line append")
	if err != nil {
		fmt.Println(err)
	}
}
