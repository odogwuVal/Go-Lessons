package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  byte
	Pin  int
}

func main() {
	p1 := Person{
		Name: "Anurag",
		Age:  22,
		Pin:  1234,
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	} else {
		// this will print a byte slice
		fmt.Println(bs)

		// this line will print the user friendly json string
		fmt.Println(string(bs))
		// As you can see filed "pin" is not there
		/*
			{"Name":"Anurag","Age":22}
		*/

	}
}
