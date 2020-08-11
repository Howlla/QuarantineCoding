package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		n,x:= reader.NextInt64(),reader.NextInt64()
		aa := make([]int64, n)
		
		for i := 0; i < len(aa); i++ {
			aa[i]=reader.NextInt64()
		}
		sol(aa,n,x)
	}
}
 
 
func sol(aa []int64, n,x int64) {
	subset:=subsets(aa,x)
	// fmt.Println(subset)
	for i:=range subset{
		sum:=int64(0)
		for j:=0;j<int(x);j++{
			sum+=subset[i][j]
		}
		if(sum%2==1){
			writer.Printf("Yes\n")
			return
		}
	}
	writer.Printf("No\n")
	return
	// numOdds:=int64(0)
	// numEven:=int64(0)
	// for i:=int64(0);i<n;i++ {
	// 	if(aa[i]%2==0){
	// 		numEven++
	// 	}else{
	// 		numOdds++
	// 	}
	// }
	// if(numOdds==0){
	// 	writer.Printf("No\n")
	// 	return
	// }
	// if(numEvens>=x-1|| numsOdds>){
	// 	writer.Printf("Yes\n")
	// }
	// if(x%2==0){

	// }

	// fmt.Println(aSum, bSum)
	// writer.Printf("%d\n",dist)

}
func subsets(nums []int64,x int64) [][]int64 {
    n := len(nums)
    var sets [][]int64;
    for mask:=0;mask<(1<<uint(n));mask++ {
        var curr []int64;
        for j:=0;j<n;j++ {
            if(mask&(1<<uint(j))!=0){
                curr=append(curr,nums[j])
            }
		}
		if(len(curr)==int(x)){
			sets=append(sets,curr);
		}
    }
    // fmt.Println(len(sets))
    return sets;
}



/*********************** I/O ***********************/


type BufferedWriter interface {
	Printf(format string, a ...interface{})
	Flush()
}
 
type writerImpl struct {
	*bufio.Writer
}
 
func NewBufferedWriter(writer io.Writer) BufferedWriter {
	return &writerImpl{Writer: bufio.NewWriter(writer)}
}
 
func (impl *writerImpl) Printf(f string, a ...interface{}) {
	fmt.Fprintf(impl.Writer, f, a...)
}
 
func (impl *writerImpl) Flush() {
	impl.Writer.Flush()
}
 
type WordScanner interface {
	NextInt() int
	NextInt64() int64
	NextString() string
}
 
type wordScannerImpl struct {
	*bufio.Scanner
}
 
func NewWordScanner(reader io.Reader) WordScanner {
	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanWords)
	// adjust the following size as needed
	// 2 << 17 is good enough to scan strings of len 2*10^5
	// it could go as high as 2 << 26 before cf judge breaks
	size := 2 << 20
	buf := make([]byte, size)
	s.Buffer(buf, size)
	return &wordScannerImpl{Scanner: s}
}
 
func (impl *wordScannerImpl) NextInt() int {
	impl.Scan()
	bb := impl.Bytes()
	i := 0
	if bb[0] == '-' {
		for _, b := range bb[1:] {
			i *= 10
			i -= int(b - '0')
		}
		return i
	}
	for _, b := range bb {
		i *= 10
		i += int(b - '0')
	}
	return i
}
 
func (impl *wordScannerImpl) NextInt64() int64 {
	impl.Scan()
	bb := impl.Bytes()
	i := int64(0)
	if bb[0] == '-' {
		for _, b := range bb[1:] {
			i *= 10
			i -= int64(b - '0')
		}
		return i
	}
	for _, b := range bb {
		i *= 10
		i += int64(b - '0')
	}
	return i
}
 

func (impl *wordScannerImpl) NextString() string {
	impl.Scan()
	return impl.Text()
}
