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
		n:= reader.NextInt64() //num of walls
		arr:=make([]int64,n)
		for i:= range arr {
			arr[i]=reader.NextInt64()
		}
		q:= reader.NextInt64();
		for ; q>0;q-- {
			x,y:= reader.NextInt64(),reader.NextInt64()
			if x==0 && y==0 {
				writer.Printf("0\n")
				continue;
			}
			ans:=0
			flag:=false;
			low:=0
			high:=len(arr)-1
			for low<=high && flag==false {
				mid:=low+(high-low)/2
				ans=directionOfPoint(x,y,arr[mid])
				if ans==-1 {
					writer.Printf("-1\n")
					flag=true
					break;
				}else if ans==0 {
					high=mid-1
				}else if ans==1 {
					low=mid+1
				}
			}
			if(flag==false){
				writer.Printf("%d\n",low)
			}
		}
	}
}
 
func directionOfPoint(x,y, p int64)int {
	x1:=int64(0)
	y1:=p

	x2:=p-x1
	y2:=0-y1
	x-=x1
	y-=y1
	crossProduct:= x2*y - y2*x
	if crossProduct>0 {
		return 1
	}else if crossProduct<0{
		return 0
	}else{
		return -1
	}
}
func sol(arr []int64,x,y int64) {
	for _,val:=range arr {
		if val%2==0 {
			writer.Printf("NO\n")
			return
		}
	}	
	writer.Printf("YES\n")
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
