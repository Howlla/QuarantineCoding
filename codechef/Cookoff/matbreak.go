package main
 

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"math"
	"math/big"
)
 


var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
	defer writer.Flush()
	
	T := reader.NextInt64()
	mod:=uint64(1000000007)	
	for t:=int64(0);t<T;t++ {

		n,a:=uint64(reader.NextInt64()),uint64(reader.NextInt64())
		sum:=uint64(a)
		curr:=uint64(a*a)
		prefixProd:=uint64(1)
		for i:=uint64(2);i<=n;i++{
			num:=uint64((i*2) -1)
			curr=((curr%mod)*(prefixProd%mod))%mod
			prefixProd=uint64(math.Pow(float64(curr),float64(num)))%mod 
			sum=((sum%mod)+(prefixProd%mod))%mod
		}
		// fmt.Println(arr)
	writer.Printf("%d \n", sum)
	}
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