package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)


func main() {
	var ch string;
	nums:=make([]int,3)
	fmt.Printf("Please enter an integer: ")
	fmt.Scan(&ch)
	fmt.Println(ch)
	for !strings.ContainsRune(ch,'x'){
	k,_:=strconv.Atoi(ch)
	nums=append(nums,k)
	nums=sort.IntSlice(nums)
	fmt.Println(nums)
	fmt.Scan(&ch)
	}
}