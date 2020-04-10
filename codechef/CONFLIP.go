package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var G int
		scanf("%d\n", &G)
		for j := 0; j < G; j++ {
			var i, n, q int
			scanf("%d %d %d\n", &i, &n, &q)
			if n%2 == 0 || i == q {
				printf("%d\n",n / 2)
			} else {
				printf("%d\n",n/2 + 1)
			}
		}
	}
}
