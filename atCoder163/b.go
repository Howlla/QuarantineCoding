package main

import (
  "fmt"
)

func main(){
  var n int64;
  fmt.Scan(&n);
  var sum int64;
  for i:=int64(1);i<=n;i++{
    if(i%3==0 || i%5==0){
    	continue;
    }else{
    	sum+=i
    }
  }
  fmt.Println(sum)

}