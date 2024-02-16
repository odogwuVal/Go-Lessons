package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	words := []string{}
	data := "A foggy mountain.\nAn old falcon.\nA wise man"

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanWords)

	n := 0
	for scanner.Scan() {
		words = append(words, scanner.Text())
		n++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# of words: %d\n", n)
	for _, word := range words {
		fmt.Println(word)
	}
}
