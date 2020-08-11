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
		aa := reader.NextString()
		bb := reader.NextString()
		sol(aa,bb,n)
	}
}
 
 
func sol(aa string, bb string, n int) {
	// for i:=0;i<n;i++{
	// 	if(aa[i]=='a'){}
	// }
	if(aa==bb){
		writer.Printf("0\n")
		return
	}
	for i:=0;i<n;i++{
		if(aa[i]=='a'&&bb[i]=='b'){
			writer.Printf("-1\n")
			return
		}
	}
	has1A,has1B,has2A,has2B:=false,false,false,false
	for i:=0;i<n;i++ {
		if(aa[i]=='a'){
			has1A=true
		}
		if(aa[i]=='b'){
			has1B=true
		}
		if(bb[i]=='a'){
			has2A=true
		}
		if(bb[i]=='b'){
			has2B=true
		}
	}
	if (has2A==true && has1A==false || has2B==true && has1B==false){
		writer.Printf("-1\n")
		return
	}
	//lets fix bs
	ans:=[][]int{}
	as:=[]int{}
	bs:=[]int{}
	flagB:=false
	for i:=0;i<n;i++ {
		if(bb[i]=='b'){
			bs=append(bs,i)
			continue
		}
		if(aa[i]=='b'&&flagB==false){
			bs=append(bs,i)
			flagB=true
			continue
		}
	}
	flag1:=false
	for _,v:=range bs {
		if(aa[v]!='b'){
			flag1=true
			break;
		}
	}
	if(flag1==false){
		bs=[]int{}
	}
	flagA:=false
	for i:=0;i<n;i++ {
		if(bb[i]=='a'){
			as=append(as,i)
			continue
		}
		if(aa[i]=='a'&&flagA==false){
			as=append(as,i)
			flagB=true
			continue;
		}
	}
	flag2:=false
	for _,val:=range as {
		if(aa[val]!='a'){
			flag2=true
			break;
		}
	}
	if(flag2==false){
		as=[]int{}
	}
	if(len(bs)!=0){
		ans=append(ans,bs)
	}
	if(len(as)!=0){
		ans=append(ans,as)
	}
	writer.Printf("%d\n",len(ans))
	for i:=0;i<len(ans);i++ {
		writer.Printf("%d ",len(ans[i]))
		for _,val:=range ans[i]{
			writer.Printf("%d ",val)
		}
		writer.Printf("\n")
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
