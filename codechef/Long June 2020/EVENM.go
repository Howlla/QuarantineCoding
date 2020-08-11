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
		n:= reader.NextInt()
		arr:=sol(n)
		for i:=0;i<n;i++ {
			for j:=0;j<n;j++ {
				writer.Printf("%d ",arr[i][j])
			}
			writer.Printf("\n")
		}
	}
}
 
 
func sol(n int) [][]int{
	row,col := 0,0
	m := n
	val:=1
	i:=0
	arr:=make([][]int,n)
	for i:= range arr {
		arr[i]=make([]int,n)
	}
	for  row < m && col < n  { 
    
        for i = col ; i < n; i++ {
			arr[ row ][i] = val
			val++
		}
        row++; 
  
       
        for i = row ; i < m ; i++ {
			arr[ i ][ n-1 ] = val
			val++ 
		}
        n--; 
  
        // Last row from remaining rows
        if  row < m  { 
            for  i = n - 1; i >= col ; i-- { 
				arr[ m-1 ][ i ] = val
				val++
			}
            m--; 
        } 
  
        if  col < n  { 
            for  i = m - 1; i >= row ; i-- {
				arr[ i ][ col ] = val
				val++
			}
            col++ ;
        } 
	}
	return arr
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
