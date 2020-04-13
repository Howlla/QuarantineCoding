package main

import (
  "fmt"
  "strings"
)

func main(){
  var n string;
  fmt.Scan(&n)
  ans:=strings.Contains(n,"7")
  if(ans==true){
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}