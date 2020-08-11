package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for ;  T>0; T-- {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		var maxA, maxB int
		for i := 0; i < n; i++ {
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			maxA = max(maxA,value)
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			maxB = max(maxB,value)
		}
		if maxA != maxB {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
