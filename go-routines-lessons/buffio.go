package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/valentinemadu/GO-ROUTINES/gofile.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buff := make([]byte, 100)

	for {
		_, err := reader.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buff))
	}
}
