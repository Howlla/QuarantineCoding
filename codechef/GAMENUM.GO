package main

import (
	"bufio"
	"fmt"
	"os"
	"math/bits"
)
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func rightRotate(n uint ,INT_BITS uint, d uint) uint { 
    /* In n>>d, first d bits are 0.  
    To put last 3 bits of at  
    first, do bitwise or of n>>d 
	with n <<(INT_BITS - d) */
	d &= 7
    return uint((n >> d)|(n << (INT_BITS - d))); 
} 
func leftRot(x uint64, k uint) uint64 {
	k &= 63
	return (x << k) | (x >> (64-k))
 }
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	scanf("%d\n", &T)
	for ;T>0;T--{
		var A,B,count uint;
		scanf("%d %d",&A,&B);
		var hm map[uint]uint
		hm = make(map[uint]uint)
		len := uint(max(bits.Len(A),bits.Len(B)))
		fmt.Println(len)
		for hm[A^B]==0 {
			hm[A^B] = count;
			count++;
			B=rightRotate(B,len,1)
			printf("%b\n",B)

		}
		var max,maxCount uint;
		
		for key, element := range hm {
			if(key>max){
				max = key
				maxCount = element
			}
		}
		printf("%d %d",maxCount,max)
	}
}