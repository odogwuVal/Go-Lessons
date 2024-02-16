package main

import "fmt"

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("only odd numbers are allowed but we got %d", i)
	}
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation Successful")
}

func main() {
	err := process(3)
	checkError(err)
	err = process(8)
	checkError(err)
}
