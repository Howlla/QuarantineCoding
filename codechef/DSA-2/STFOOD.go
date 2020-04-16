package main

import (
	"bufio"
	"fmt"
	"os"

)

var bytes []byte
var l, maxB int


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var T int;
    scanner.Scan()

    fmt.Sscanf(scanner.Text(), "%d", &T)
	for ;T>0;T--{
		var n int;
		scanner.Scan()

    	fmt.Sscanf(scanner.Text(), "%d", &n)
		var attemptedA,attemptedB, A,B,i,ans int;
		ans=2*n;
		scanner.Scan()
		var arr string;
    	fmt.Sscanf(scanner.Text(), "%s", &arr)
    	fmt.Println(arr)
			for ;i<2*n;{
			if(i%2==0){
				if(arr[i]=='1'){
					A++
				}
				attemptedA++
			}else{
				if(arr[i]=='1'){
					B++
				}
				attemptedB++
			}
			if((A-B)>(n-attemptedB) || (B-A)>(n-attemptedA)){
				ans=i+1
				break;
			}
			i++;
		}
		fmt.Println(ans)
	}

}
