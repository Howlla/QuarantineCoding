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
	for j := 0; j < T; j++ {
		var n int
		var count int
		scanf("%d\n", &n)
		for i := 5; n/i > 0; i *= 5 {
			count += n / i
		}
		printf("%d\n",count)
	}
}
