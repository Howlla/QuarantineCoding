package main

import (
	"bufio"
	"fmt"
	"os"

)
func max(a,b int) int{
	if(a>b){
		return a
	}
	return b;
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int;
	fmt.Fscanln(reader, &T)
	for t:=0;t<T;t++ {
		var n,k int;
		fmt.Fscanln(reader,&n,&k)
		countsEveryItteration := n-1
		multiples := k/countsEveryItteration
		if(k%countsEveryItteration==0){
			fmt.Fprintln(writer,n*multiples-1)
			continue
		}
		startValue := n*multiples
		startItterations := countsEveryItteration*multiples
		diff:= k -startItterations
		fmt.Fprintln(writer,startValue+diff)
	}
}