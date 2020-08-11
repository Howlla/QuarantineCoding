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
		// n,k:= reader.NextInt64(), reader.NextInt64()
		a,b,c,p,q,r:=reader.NextInt(), reader.NextInt(),reader.NextInt(), reader.NextInt(),reader.NextInt(), reader.NextInt()
		writer.Printf("%d\n",sol(a,b,c,p,q,r))
	}
}
 
 
func sol(a,b,c,p,q,r int) int {
	same:=0
	if(a==p){
		same++
	}
	if(b==q){
		same++
	}
	if(c==r){
		same++
	}
	
	if(same==3){
		return 0
	}else if(same==2){
		return 1
	}else if(same==1){
		//check which one is same
		// x1,x2,y1,y2:=0,0,0,0
		if(a==p){
			if((q-b==r-c) || (b!=0 && c!=0 && q%b==0 && r%c==0 && (q/b)==(r/c))){
               return 1
            }
		}
		if(b==q){
			if((p-a==r-c) || (a!=0 && c!=0 && p%a==0 && r%c==0 && (p/a)==(r/c))){
                return 1 
            }
		}
		if(c==r){
			if((p-a==q-b) || (a!=0 && b!=0 && p%a==0 && q%b==0 && (p/a)==(q/b))){
               return 1
            }
		}
		//check if difference between them is same
		return 2
	}else{
		//No similar numbers same==0
		if((p-a==q-b && q-b==r-c) || (b!=0 && c!=0 && a!=0 && p%a==0 && q%b==0 && r%c==0 && (p/a)==(q/b) && (q/b)==(r/c))){
            return 1
        }
        if((q-b==r-c) || (b!=0 && c!=0 && q%b==0 && r%c==0 && (q/b)==(r/c))){
           return 2
        }
        if((p-a==r-c) || (a!=0 && c!=0 && p%a==0 && r%c==0 && (p/a)==(r/c))){
            return 2
        }
        if((p-a==q-b) || (a!=0 && b!=0 && p%a==0 && q%b==0 && (p/a)==(q/b))) {
            return 2
        }
        if((a-p)+(b-q)==(c-r) || (b-q)+(c-r)==(a-p) || (a-p)+(c-r)==(b-q)){
            return 2
        }
        if(p!=0 && q!=0 && r!=0 && a%p==0 && b%q==0 && c%r==0 && ((a/p)*(b/q)==(c/r) || (a/p)*(c/r)==(b/q) || (b/q)*(c/r)==(a/p))){
           return 2
        }
        if(p!=0 && a%p==0 && ((b-((a/p)*q))==c-r || (c-((a/p)*r))==b-q)) {
            return 2
        }
        if(q!=0 && b%q==0 && ((a-((b/q)*p))==c-r || (c-((b/q)*r))==a-p)){
          return 2
        }
        if(r!=0 && c%r==0 && ((b-((c/r)*q))==a-p || (a-((c/r)*p))==b-q)) {
            return 2
        }
        if(r!=0 && c%r==0 && ((q+a-p!=0 && b%(q+a-p)==0 && b/(q+a-p)==c/r) || (p+b-q!=0 && a%(p+b-q)==0 && a/(p+b-q)==c/r))) {
            return 2
        }
        if(q!=0 && b%q==0 && ((p+c-r!=0 && a%(p+c-r)==0 && a/(p+c-r)==b/q) || (r+a-p!=0 && c%(r+a-p)==0 && c/(r+a-p)==b/q))) {
            return 2
        }
        if(p!=0 && a%p==0 && ((q+c-r!=0 && b%(q+c-r)==0 && b/(q+c-r)==a/p) || (r+b-q!=0 && c%(r+b-q)==0 && c/(r+b-q)==a/p))){
           return 2
        }
	}
	return 3
}


func min(nums ...int) int {
    ans := nums[0]
    for _, num := range nums {
        if num<ans{
			ans=num
		}
    }
    return ans
}
func abs(a int) int {
	if(a<0){
		return -a
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
