package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 
var mod int64 = 1000000007
 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
	// t := reader.NextInt64()
	aa:=make([]int64,100000007)
	calc(aa)
	for ; t > 0; t-- {
		l:= reader.NextInt64()
		r := reader.NextInt64()
		ans:=int64(0)
		for i := l; i <= r; i++ {
			ans+=((aa[i]%mod)*(aa[i]%mod))%mod
		}
		writer.Printf("%d\n",ans%mod)
	}
	fmt.Println(aa[100000000:100000007])
}
 
func calc(aa []int64) {
	aa[1] = 1
	for i:=2;i<100000007;i++ {
		aa[i] = 1 + aa[i-int(aa[aa[i-1]])]%mod
	}
}
 
func sol(aa []int,bb []int, n int) {

	writer.Printf("%d\n",1)
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
