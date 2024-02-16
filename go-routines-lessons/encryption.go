package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodedWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password has been encrypted to: ", encodedWithMD5(password))
}
