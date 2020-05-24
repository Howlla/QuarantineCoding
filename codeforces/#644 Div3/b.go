package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
  "strings"
  "strconv"
  "math"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		n:= reader.NextInt64()
		aa := make([]int, n)
		for i := 0; i < len(aa); i++ {
			aa[i] = reader.NextInt()
		}
		sort.Ints(aa)
		writer.Printf("%d\n", sol(aa))
	}
}
 
func sol(aa []int) int {
	minDiff:=aa[1]-aa[0]
	for i:=1;i<len(aa)-1;i++ {
		if(aa[i+1]-aa[i] < minDiff){
			minDiff = aa[i+1]-aa[i]
		}
	}
	return minDiff
  
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
func SArr(a []int) string {
	sb := strings.Builder{}
	sb.Grow(len(a) * 9)
	for i := range a {
		sb.WriteString(strconv.Itoa(a[i]) + " ")
	}
	return sb.String()
}
func SArr64(a []int64) string {
	sb := strings.Builder{}
	sb.Grow(len(a) * 9)
	for i := range a {
		sb.WriteString(strconv.FormatInt(a[i], 10) + " ")
	}
	return sb.String()
}
func MnMx(args ...int64) (int64, int64) {
	minVal, maxVal := int64(math.MaxInt64), int64(math.MinInt64)
	for _, v := range args {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return minVal, maxVal
}
func Mn(args ...int64) int64 { mn, _ := MnMx(args...); return mn }
func Mx(args ...int64) int64 { _, mx := MnMx(args...); return mx }
func LCM(a, b int64) int64   { return a / GCD(a, b) * b }
func GCD(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}
func sortInt64Slice(a []int64) { sort.Slice(a, func(i, j int) bool { return a[i] < a[j] }) }
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}