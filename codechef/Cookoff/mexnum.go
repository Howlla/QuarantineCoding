package main
 

import (
	"bufio"
	"fmt"
	"os"
	"io"
)
 


var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
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
func subsets(nums []int64) [][]int64 {
    n := len(nums)
    var sets [][]int64;
    for mask:=0;mask<(1<<uint(n));mask++ {
        var curr []int64;
        for j:=0;j<len(nums);j++ {
            if(mask&(1<<uint(j))!=0){
                curr=append(curr,nums[j])
            }
        }
         sets=append(sets,curr);
    }
    return sets;
}
func firstMissingPositive(nums []int64) int64 {
    hashmap := make(map[int64]int64)
    for i:=int64(0); i<int64(len(nums));i++ {
        hashmap[nums[i]] = i
    }
    for i:=int64(1); i<=int64(len(nums));i++ {
        if _ , ok := hashmap[i]; !ok {
            return i
        }
    }
    return int64(len(nums)+1)
}



func main() {
	defer writer.Flush()
	
	T := reader.NextInt64()
	mod:=int64(998244353)	
	for t:=int64(0);t<T;t++ {
		n:=(reader.NextInt64())
		arr:=make([]int64,n)
		for i:=int64(0);i<n;i++{
			arr[i]=(reader.NextInt64())
		}
		allSubsets:=subsets(arr)
		s:=int64(0);
		for i:=0;i<len(allSubsets);i++{
			s=(s+(firstMissingPositive(allSubsets[i])%mod))%mod
		}
		s=s%mod
		writer.Printf("%d \n", s%mod)
	}
}