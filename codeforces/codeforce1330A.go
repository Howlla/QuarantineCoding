package main

import (
	// "bufio"
	"fmt"
	// "strings"
	// "strconv"
	// "os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	var T int
	fmt.Scan(&T)
	for ; T>0; T-- {
		solve()
	}
}
func solve() {
	var n,x int;
		fmt.Scan(&n,&x);
		hm:=make(map[int]bool);
		arr:=make([]int,n)
		for i:=0;i<n;i++ {
			fmt.Scan(&arr[i])
			hm[arr[i]]=true
		}
		for i:=n+x;i>0;i--{
			var v int;
			for j:=1;j<=i;j++ {
				if(hm[j]==false){
					v++
				}
			}
			if(v<=x){
				fmt.Printf("%d \n",i)
				return;
			}
		}
}