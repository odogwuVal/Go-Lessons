package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("learning go seems fun")
	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]))
		if err != nil {
			fmt.Println("breaking")
			break
		}
	}

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
