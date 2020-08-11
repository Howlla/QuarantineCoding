package main
 

import (
	"bufio"
	"fmt"
	"os"

)
 

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var T int;
	fmt.Fscanln(reader, &T)
	for t:=0;t<T;t++ {
		var n,q int64;
		fmt.Fscanln(reader,&n,&q)
		sum:=int64(0)
		curr:=int64(0)
		for i:=int64(0);i<q;i++{
			var f,d int64
			fmt.Fscanln(reader,&f,&d)
			if(f>n||d>n){
				return
			}
			if(i==0){
				sum+=f
				sum+=abs(f-d)
				curr=d
			}else{
				sum+=abs(curr-f)
				sum+=abs(f-d)
				curr=d
			}
		}
		fmt.Fprintln(writer,sum)
	}
}
 

func abs(a int64)int64{
	if(a>=0){
		return a
	}
	return -a
}

