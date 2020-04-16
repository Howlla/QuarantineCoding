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
    
    //take input
    for i:=0;i<n;i++{
      if i == n-1 {
        fmt.Fscanf(reader, "%d\n", &nums[i])
      } else {
        fmt.Fscanf(reader, "%d", &nums[i])
      }   
    }
    var sumTotal int;
    for nums[0]!=0 {
      var indexMin,min int;
      min = 9999999
      // find index of minimum value
      for i:=0;i<n;i++{
        if nums[i]<min && nums[i]!=0{
          min=nums[i]
          indexMin = i
        }
      }
      sumTotal+= n*min
      for i:=0;i<n;i++{
        nums[i]=nums[i]-min
      }
      n=indexMin
    }
   fmt.Println(sumTotal)
  }
}


func min(a,b int) int {
  if(a>b){
    return b
  }
  return a
}