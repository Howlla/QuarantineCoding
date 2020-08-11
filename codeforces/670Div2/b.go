package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
	"strconv"
	"strings"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
var dp []string; 

func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	dp=make([]string,100005)
	calculate(dp)
	for ; t > 0; t-- {
		n:= reader.NextInt64()
		// writer.Printf("%s\n",sol(n))
		writer.Printf("%s\n",sol_eff(n))
	}
}

func calculate (dp[]string) {
	for i:=1;i<=10005;i++ {
		if i==1 {
			dp[i] = "8"
			continue
		}
		if(i%4==1) {
			dp[i] = (dp[i-1] + 8)*10
		}else if(i%4==2 || i%4==3) {
			dp[i] = 9 * int(math.Pow(float64(10),float64(i-1))) + dp[i-1]
		}else {
			dp[i] = (dp[i-1]*10)+10
		}
	}
}
func sol_eff(n int64) string {
	return dp[n]
}
 
func sol(n int64) string {
	// str := (strconv.FormatInt(n, 2))

	aa :=make([]string,n);
	for i:=0;i<int(n);i++ {
		aa[i] = "1001"
	}
	num := strings.Join(aa,"")
	y,_:=strconv.ParseInt(num,2,64);
	multiply :=(strconv.FormatInt((math.MaxInt64),2))
	// fmt.Printf("%T",multiply)
	runeArr := []byte(multiply)
	for i:=len(runeArr)-1;i>=len(runeArr)-int(n);i-- {
		runeArr[i] = '0';
	}
	multiply = string(runeArr)
	// fmt.Println(multiply)
	x,_ := strconv.ParseInt(multiply, 2, 64);
	// num:=int64(0)
	// for i:=0;i<int(n);i++ {
	// 	num = num*10 + 9
	// }
	ansInt:= y&x
	ansString := strconv.FormatInt(ansInt,2)
	bb:=make([]string,n)
	// fmt.Println(ansString)
	for i:=0;i<int(n);i++ {

		z,_:= strconv.ParseInt(ansString[i*4:i*4+4],2,64)
		bb[i]=strconv.Itoa(int(z));
	}
	retValue := ""
	for _,v:=range bb{
		retValue+=v
	}
	return retValue
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
