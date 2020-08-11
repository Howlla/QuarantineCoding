package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 

 func max(a,b int)int{
	 if(a>b){
		 return a
	 }
	 return b
 }
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 var s string;
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		s= reader.NextString()
		q:=reader.NextInt64()
		n:=len(s)
		pos:=make([]int,len(s))
		index=n
		for i:=n-1;i>=0;i-- {
			if s[i]=='(' {
				index=i
			}
			pos[i]=index
		}
		for ;q>0;q-- {
			index:=reader.NextInt64()
			start:=pos[index-1]
		}
	}
}
 
 
func isBalanced(i,mid int64) bool{
			counter:=1
			i++
			for i<=mid{
				if s[i]=='('{
					counter++
				}else{
					if counter>0{
						counter--
						if(counter==0){
							return true
						}
					}
				}
				i++
			}
		return false
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
