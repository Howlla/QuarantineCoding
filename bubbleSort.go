package main
import "fmt"

func main(){
	array:=make([]int,10)
	for i:=0;i<10;i++{
		fmt.Scan(&array[i])
	}
	BubbleSort(array)
	fmt.Println(array)
}

func BubbleSort(arr []int){
	n:=len(arr)
	for i:=0;i<n;i++{
		swapped:=false
			for j:=0;j<n-i-1;j++{
				if(arr[j]>arr[j+1]){
					swap(arr,j)
					swapped=true
				}
			}
		if(swapped==false){
			break;
		}
	}
}

func swap(s []int, x int) {
	s[x], s[x+1] = s[x+1], s[x]
}