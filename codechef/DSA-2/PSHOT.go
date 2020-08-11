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
	t := reader.NextInt()
	for ; t > 0; t-- {
		n:= reader.NextInt()
		// aa := make([]string, 2*n)
		// for i := 0; i < len(aa); i++ {
		// 	aa[i] = reader.NextString()
		// }
		aa:=reader.NextString()
		writer.Printf("%d\n", sol(aa, n))
	}
}
 
func sol(aa string, n int) int {
	var a,b,aScore,bScore int;
	for i:= range(aa) {
		if(i%2==0){
			if(aa[i]=='1'){
				aScore++
			}
			a++
		}else{
			if(aa[i]=='1'){
				bScore++
			}
			b++
			
		}
		if((aScore-bScore)>n-b || (bScore-aScore)>n-a){
						fmt.Pritnln(aScore,bScore)

			return i+1
		}
		
	}
	return 2*n;
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
/*********************** Utils ***********************/
 
func S(v interface{}) string { return fmt.Sprintf("%v", v) }


func LCM(a, b int64) int64   { return a / GCD(a, b) * b }
func GCD(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}