package main

import "fmt"

func main(){
	var input string;
	fmt.Scan(&input);
	var flag bool;
	for i:=0;i<len(input);i++ {
		if input[i]=="i"||input[i]=="a"||input[i]=="n"{
			flag=true
		}
	}
	if(flag==true){
		fmt.Println("Found!")
	}else{
		fmt.Println("Not Found!")
	}
	
}