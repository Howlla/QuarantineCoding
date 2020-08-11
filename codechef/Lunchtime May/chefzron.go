package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		n:= reader.NextInt()
		aa := make([]int, n)
		bb := make([]int, 2*n)
		for i := 0; i < len(aa); i++ {
			aa[i]=reader.NextInt()
			bb[i]=aa[i]
			bb[n+i]=aa[i]
		}
		sol(aa,bb,n)
	}
}
 
 
func sol(aa []int,bb []int, n int) {
	counts:=0
	for i:=0; i<n ;i++ {
		if(aa[i]==1){
			counts++
		}
	}
	if(counts==0){
		writer.Printf("0 \n")
		return
	}
	if(counts==1 ){
		writer.Printf("-1 \n")
		return
	}
	maxDist:=math.MaxInt32
	for i:=0;i<n;i++ {
		hm:=make(map[int]int)
		max1:=0
		for j:=i ;j<i+n; j++ {
			if(bb[j]==1){
				if _,found:= hm[bb[j]];!found{
					hm[bb[j]]=j
				}else{
					max1=max(max1,j-hm[1])
				}
			}
		}
		maxDist = min(maxDist,max1)
	}
	writer.Printf("%d\n",maxDist)
}

func max (a,b int) int{
	if(a>b){
		return a
	}
	return b
}

func min (a,b int) int {
	if(a>b){
		return b
	}
	return a
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
