package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
// var adjacencyList map[int][]int;
// var currSet map[int]bool;
// var maxSet map[int]bool;
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		n,m:= reader.NextInt(),reader.NextInt()
		income:=make([]int64,n+1)
		population:=make([]int64,n+1)
		for i:=1;i<n+1;i++{
			income[i]=reader.NextInt64()
		}
		for i:=1;i<n+1;i++{
			population[i]=reader.NextInt64()
		}
		g:=make([][]int64,200005)
		for i:= range g {
			g[i]=make([]int64,0)
		}

		for i:=0;i<m;i++ {
			u:=reader.NextInt64()
			v:=reader.NextInt64()
			g[u]=append(g[u],v)
			g[v]=append(g[v],u)
		}
		num:=income[1]
		den:=population[1]
		for i:=2;i<n;i++ {
			if(income[i]*den > population[i]*num){
				den=population[i];
				num=income[i];
			}
		}
		v:=make([]int64,n+1)
		ans:=[]int64{}
		for i:=1;i<n+1;i++ {
			if(v[i]==0 && num*population[i]==den*income[i]) {
				queue:=[]int64{int64(i)}
				temp:=[]int64{}
				v[i]=1
				for len(queue)!=0 {
					top:=queue[0]
					queue=queue[1:]
					temp=append(temp,top)
					for _,x:=range g[top] {
						if(v[x]==0 && num*population[x]==den*income[x]){
							queue=append(queue,x)
							v[x]=1
						}
					}
				}
				if len(temp)>len(ans){
					ans=temp
				}
			}
		}
		writer.Printf("%d\n",len(ans))
		for _,k:=range ans {
			writer.Printf("%d ",k)
		}
		writer.Printf("\n");
	}
}
 

// func dfs (node int,visited *[]bool){
// 	(*visited)[node]=true
// 	currSet[node+1]=true
// 	for _,v:=range adjacencyList[node] {
// 		if((*visited)[v]==false){
// 			dfs(v,visited)
// 		}
// 	}
// }
// func connectedComponents(n int) {
// 	visited:=make([]bool,n)
// 	for v:=0;v<n;v++ {
// 		if(visited[v]==false) {
// 			currSet=make(map[int]bool)
// 			dfs(v,&visited);
// 			if len(currSet)>len(maxSet) {
// 				maxSet=currSet
// 			}
// 		}
// 	}
// }	


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
