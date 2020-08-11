package main
 

import (
	"bufio"
	"fmt"
	"os"
	"io"
)
 


var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)

func max(a,b int) int{
	if (a>b){
		return a
	}
	return b
}

func abs(a int64)int64{
	if(a>=0){
		return a
	}
	return -a
}

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


func main() {
	defer writer.Flush()
		n:=(reader.NextInt())
		arr:=make([]int,n)
		for i:=0;i<n;i++{
			arr[i]=(reader.NextInt())
		}
		cache:=make([][]int,n)
		for i:=range cache{
			cache[i]=make([]int,n)
		}

		writer.Printf("%d \n", solve(arr,0,n-1,n,cache))
}

func solve(wines []int,be int, end int,n int,cache [][]int) int {
	if ( be > end ) {
		return 0
	}
	year:= n-(end-be)
	if(cache[be][end]!=0){
		return cache[be][end]
	}
	cache[be][end] = max(
		solve(wines,be+1, end,n,cache)+wines[be]*year,
		solve(wines,be,end-1,n,cache)+wines[end]*year)
	return cache[be][end]
}