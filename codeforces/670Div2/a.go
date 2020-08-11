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
		n:= reader.NextInt64()
		// // s:=reader.NextString();
		// k:= reader.NextInt64()
		// aa := make([]int64, n)
		// for i := 0; i < len(aa); i++ {
		// 	aa[i]=reader.NextInt64()
		// }
		sol(n)
	}
}
 
 
func sol(n int64) {
	if n<=30 {
		writer.Printf("NO\n")
	}else{
		writer.Printf("YES\n")
		if (n==36){
			writer.Printf("5 6 10 15\n")
		}else if (n==40) {
			writer.Printf("6 10 15 9\n")
		}else if (n==44){
			writer.Printf("6 7 10 21\n")
		}else{
			writer.Printf("6 10 14 %d\n",n-30)
		}
	}
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
