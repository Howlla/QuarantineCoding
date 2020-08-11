package main

import (
	"fmt"

)

func main(){
  var x float64;
  fmt.Scan(&x);
  principal:=float64(100)
  year:=float64(0)
  for total:=principal;total<x;year++{
	total*=1.01
	total=float64(int(total))
  }
  fmt.Println(year)
}