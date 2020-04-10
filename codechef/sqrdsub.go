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
		var countGood,prod,count int;
		prod=1;
		var preProd,countZero []int;
		preProd=make([]int,n)
		countZero = make([]int,n)
		for i:=0;i<n;i++ {
			if(nums[i]==0){
				count++
			}else{
				prod*=nums[i]
			}
			preProd[i]=prod
			countZero[i]=count
		}
		for i:=0;i<n;i++{
			for j:=i;j<n;j++{
				if(i==0){
					if(countZero[j]==0){
						prod=preProd[j]
					}else{
						prod=0
					}
				}else{
					if(countZero[j]-countZero[i-1]==0){
					prod=preProd[j]/preProd[i-1];
					}else{
						prod=0
					}
				}
				if((prod%4==0||prod%2!=0) && prod!=0){
					// fmt.Println(prod)
					countGood++
				}
			}
		}
		fmt.Println(countGood)
	}
}