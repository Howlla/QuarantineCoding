package main

import (
	"fmt"
	"strings"
)

func main() {
	z := new(string)
	fmt.Printf("Please enter string: ")
	fmt.Scan(z)
	s := strings.ToLower(*z)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
