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
		var n int;
		fmt.Fscanln(reader,&n)
		nums:=make([]int,n)
		
		for i:=0;i<n;i++{
			if i == n-1 {
				fmt.Fscanf(reader, "%d\n", &nums[i])
				
			} else {
				fmt.Fscanf(reader, "%d", &nums[i])
			}	
		}
	flag:=true
	i:=0;
	j:=n-1;
	lastMove:=0
	moves:=0
	a:=0
	b:=0
	for(i<=j){
		moves++
		if(flag==true){
			curr:=0;
			for i<=j && curr<=lastMove {
				curr+=nums[i]
				i++
			}
			lastMove=curr
			a+=curr
		}else{
			curr:=0;
			for i<=j && curr<=lastMove {
				curr+=nums[j]
				j--
			}
			lastMove=curr
			b+=curr
		}
		flag=!flag
	}


		fmt.Fprintln(writer,moves,a,b);
	}
}