package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "fmt"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
func main() {
	defer writer.Flush()
		n:= reader.NextInt()
		k:= reader.NextInt()
		dp:=make([][]int,n+1)
		for i:=range dp {
			dp[i]=make([]int,k+1)
		}
		fmt.Println(solve_dp(n,k,dp))
}
func solve_dp(n,k int,dp [][]int) int {
	K:= k
	N:= n
	for n=0;n<=N; n++ {
		for k = 0;k<=K;k++ {
			if k==1 || k==n {
				dp[n][k] = 1
				continue
			}
			if k<=0 || n<=0 || k>n {
				dp[n][k]=0
				continue
			} 
			dp[n][k] =  dp[n-1][k]* k + dp[n-1][k-1]
		}
	}
	return dp[N][K];
}
 
func solve(n,k int, dp [][]int) int{
	if k==1 || k==n {
		dp[n][k] = 1
		return 1
	}
	if k<=0 || n<=0 || k>n {
		dp[n][k]=0
		return 0
	} 
	if dp[n][k]!=0{
		return dp[n][k]
	}
	dp[n][k] =  solve(n-1,k,dp) * k + solve(n-1,k-1,dp)
	return dp[n][k]
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
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
