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
		if(n%2==1&&k%2==0){
			fmt.Fprintln(writer,"NO");
			continue
		}
		if(n%2==0 &&k%2==1){
			if(n<2*k){
			fmt.Fprintln(writer,"NO");
			continue
			}
			fmt.Fprintln(writer,"YES");

			for i:=1;i<k;i++{
				fmt.Fprint(writer,"2 ")
			}
			fmt.Fprintln(writer,n-2*(k-1))
		}else{
			if(n<k){
				fmt.Fprintln(writer,"NO");
				continue
			}
			fmt.Fprintln(writer,"YES");
			for i:=1;i<k;i++{
				fmt.Fprint(writer,"1 ")
			}
			fmt.Fprintln(writer,n-k+1)
		}
	}
}

