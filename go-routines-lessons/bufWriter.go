package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/Users/valentinemadu/GO-ROUTINES/bufio.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

}
