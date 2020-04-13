package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	solveNsquare(k)
	// var ans int
	// for i := 1; i <= k; i++ {
	// 	for j := 1; j <= k; j++ {
	// 		for p := 1; p <= k; p++ {
	// 			 gcdij := gcd(i, j)
	// 			 gcdijk := gcd(gcdij, p)
	// 			ans += gcdijk
	// 		}
	// 	}
	// }
	// fmt.Println(ans)
}

func solveNsquare(k int){
 mp:=make(map[int]int)
 var sum int;
 for i:=1;i<=k;i++{
	 for j:=1;j<=k;j++{
		 mp[i]+=gcd(i,j)
	 }
 }
 for i:=1;i<=k;i++{
	for j:=1;j<=k;j++{
		g:=gcd(i,j)
		sum+=mp[g]
	}
 }
 fmt.Println(sum)
}

func gcd(a, b int) int{
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
