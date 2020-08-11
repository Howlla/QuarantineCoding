package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

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
		var x int64;
		fmt.Fscanln(reader,&n,&x)
		nums:=make([]int,n)
		var sum int64;
		sum=0
		// count:=0
		for i:=0;i<n;i++{
			if i == n-1 {
				fmt.Fscanf(reader, "%d\n", &nums[i])
				
			} else {
				fmt.Fscanf(reader, "%d", &nums[i])
			}	
		
		}
		// sum = sum - (count*x)
		sort.Ints(nums);
		// fmt.Println(sum,nums,count)
		// for i:=len(nums)-1-count;i>=0;i--{
		// 	diff:= x-nums[i];
		// 	if(sum-diff>=0){
		// 	sum-=diff;

		// 	count++
		// }
		// }	
		prefix:=n;
		for prefix>0 && sum+int64(nums[prefix-1])-x >=0 {
				prefix--
				sum= sum+int64(nums[prefix])-x
		}



		// for i:=len(nums)-count-1;i>=0;i--{
		// 	fmt.Println(nums[i],"here")
		// 	if( sum+nums[i]>=(count+1)*x ){
		// 		count++
		// 		sum+=nums[i]
		// 	}else{
		// 		break;
		// 	}
		// }






		fmt.Fprintln(writer,n-prefix);
	}
}