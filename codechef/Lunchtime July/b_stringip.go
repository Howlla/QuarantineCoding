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
		m:=reader.NextInt64();
		low:=reader.NextInt64();
		high:=reader.NextInt64();
		if (n==1 && m==1 ){
			writer.Printf("%d\n",low);
			continue;
		}
		if low>high {
			low = high
		}
		if (2*low<=high) {
			writer.Printf("%d\n",low*n*m);
			continue;
		}
		dp:=make([][]int64,n)
		for i:=range dp {
			dp[i]=make([]int64,m)
		}
		sum:=int64(0)
		for i:=0;i<int(n);i++ {
			for j:=0;j<int(m);j++ {
				if (i%2==0){
					if j%2==0 {
						dp[i][j] = low
						sum += low
					}else{
						dp[i][j] = high-low
						sum+= high - low
					}
				}else{
					if j%2==0 {
						dp[i][j] = high-low
						sum += high-low
					}else{
						dp[i][j] = low
						sum+= low
					}
				}
			}
		}
		// for i:=range dp{
		// 	fmt.Println();
		// 	for j:=range dp[0] {
		// 		fmt.Print(dp[i][j])
		// 	}
		// 	fmt.Println()
		// }
		writer.Printf("%d\n",sum);
	
	}
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
