package main

import (
	"bufio"
	"fmt"
	"os"
	"math"

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
		var n int;
		fmt.Fscanln(reader,&n)
		k:=0;
		count:=0
		temp:=n
		ans:=[]int{}
		for temp!=0 {
			if(temp%10!=0){
			ans=append(ans,(temp%10)* int(math.Pow(10.0,float64(count))))
			k++
			}
			count++
			temp=temp/10
		}
		fmt.Fprintln(writer,k);
		for _,val := range ans {
			fmt.Fprint(writer,val," ");
		}
		fmt.Fprintln(writer,"")
	}
}