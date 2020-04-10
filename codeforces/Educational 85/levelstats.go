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
		fmt.Fscanln(reader,&n);
		var attempted,cleared int;
		var prevAttempted,prevCleared int;
		flag:= false
		for i:=0;i<n;i++{
			fmt.Fscanln(reader, &attempted,&cleared)
			if(i==0){
				prevAttempted=attempted
				prevCleared=cleared
			}
			cDiff:=cleared-prevCleared
			aDiff:=attempted-prevAttempted
			if(cleared>attempted||prevAttempted>attempted||prevCleared>cleared ||cDiff>aDiff){
				flag=true
			}else{
				prevAttempted=attempted
				prevCleared=cleared
			}
		}
		if(flag==false){
			fmt.Fprintln(writer,"YES")
		}else{
			fmt.Fprintln(writer,"NO")
		}
	}
}