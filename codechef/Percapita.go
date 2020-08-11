package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 

 
var reader = NewWordScanner(os.Stdin)
var writer = NewBufferedWriter(os.Stdout)
 
var adjacencyList map[int][]int;
var currSet map[int]bool;
var maxSet map[int]bool;
func main() {
	defer writer.Flush()
	t := reader.NextInt64()
	for ; t > 0; t-- {
		n,m:= reader.NextInt(),reader.NextInt()
		income:=make([]int,n)
		population:=make([]int,n)
		perCapita:=make([]int,n)
		for i:=0;i<n;i++{
			x:=reader.NextInt()
			income[i]=x
			perCapita[i]=x
		}
		max:=0
		for i:=0;i<n;i++{
			x:=reader.NextInt()
			population[i]=x
			perCapita[i] /= x
			if(perCapita[i]>max){
				max=perCapita[i]
			}
		}
		adjacencyList=make(map[int][]int)
		citiesMax:=make(map[int]bool)
		for i:=0;i<n;i++ {
			if perCapita[i]==max {
				citiesMax[i+1]=true
			}
			adjacencyList[i]=make([]int,0)

		}
		maxSet=make(map[int]bool)
		for i:=0;i<m;i++ {
			u:=reader.NextInt()
			v:=reader.NextInt()
			if(citiesMax[u]==true && citiesMax[v]==true){
				addEdge(u-1,v-1)
			}
		}
		connectedComponents(n)
		writer.Printf("%d\n",len(maxSet))
		for k:=range maxSet {
			writer.Printf("%d ",k)
		}
		writer.Printf("\n");
	}
}
 

func addEdge(m,n int) {
	adjacencyList[m]=append(adjacencyList[m],n)
	adjacencyList[n]=append(adjacencyList[n],m)
}
func dfs (node int,visited *[]bool){
	(*visited)[node]=true
	currSet[node+1]=true
	for _,v:=range adjacencyList[node] {
		if((*visited)[v]==false){
			dfs(v,visited)
		}
	}
}
func connectedComponents(n int) {
	visited:=make([]bool,n)
	for v:=0;v<n;v++ {
		if(visited[v]==false) {
			currSet=make(map[int]bool)
			dfs(v,&visited);
			if len(currSet)>len(maxSet) {
				maxSet=currSet
			}
		}
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
