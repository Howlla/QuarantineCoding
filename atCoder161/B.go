package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"math"
)

func min(a,b int)int{
	if(a>b){
		return b
	}
	return a
}

func max(a,b int)int{
	if(a>b){
		return a
	}
	return b
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()
	var N,M int;
	fmt.Scan(&N,&M);

	var total int;
	arr := make([]int, N)
		for i := range arr {
			fmt.Scan(&arr[i])
			total+=arr[i]
		}
	var threshold,count int;
	threshold = int(math.Ceil(float64(total)/float64((4*M))));
	sort.Ints(arr);
		for i:=0;i<N;i++{
			if(arr[i]>=threshold){
				count++
			}
		}
		if(count>=M){
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	
}
