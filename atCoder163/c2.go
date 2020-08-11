package main

import (
	"bufio"
	"fmt"
	"os"

)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
  var n int64;
  fmt.Scan(&n);
  nums:=make([]int64,n)

  for i:=1;i<int(n);i++{
			if i == int(n-1) {
				fmt.Fscanf(reader, "%d\n", &nums[i])
			} else {
				fmt.Fscanf(reader, "%d", &nums[i])
			}		
		}
	fmt.Println(nums)
  ans:=make([]int64,n)
  for i:=1;i<int(len(nums));i++{

	  ans[nums[i]-1]++
  }
 for i:=range ans{
	 fmt.Println(ans[i])
 }
}