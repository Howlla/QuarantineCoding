//Day 1
func subarraySum(nums []int, k int) int {
    ans:=0
    hm:=make(map[int]int);
    hm[0]=1
    currSum:=0
    for i:=0;i<len(nums);i++ {
        currSum+=nums[i]
        ans+=hm[currSum-k]
        hm[currSum]++
    }
    return ans
}

//Day 2
func rangeBitwiseAnd(m int, n int) int {
    count:=0
    for m!=n{
        m=m>>1
        n=n>>1
        count++
    }
    return m<<count
}

//Day 3
import "container/list"

type LRUCache struct {
    cache map[int]*list.Element
    capacity int
    ll *list.List
}

type node struct{
    key int
    value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
        ll:list.New(),
        capacity:capacity,
        cache:make(map[int]*list.Element,capacity),
    }
}


func (this *LRUCache) Get(key int) int {
    if val,ok:=this.cache[key];ok{
        this.ll.MoveToFront(val)
        return val.Value.(*node).value
    }else{
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if val,ok:=this.cache[key];ok {
        this.ll.MoveToFront(val)
        val.Value= &node{key,value}
    }else{
        if(this.ll.Len()==this.capacity){
            nodeToRemove:=this.ll.Back()
            keyToRemove:=nodeToRemove.Value.(*node).key
            delete(this.cache,keyToRemove)
            this.ll.Remove(nodeToRemove)
        }
        this.ll.PushFront(&node{key,value})
        this.cache[key]=this.ll.Front()
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

 //Day 4

 func canJump(nums []int) bool {
    if(len(nums)==0){
        return true
    }
    farthestJump:=0
    n:=len(nums)
    for i:=0;i<n;i++{
        if i>farthestJump{
            return false
        }else{
            farthestJump=max(farthestJump,i+nums[i])
        }
    }
    return true
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

//Day 5

func longestCommonSubsequence(text1 string, text2 string) int {
    rows:= len(text1)+1
    cols:=len(text2)+1
    dp:= make([][]int, rows)       // initialize a slice of dy slices
    for i:=0;i<rows;i++ {
        dp[i] = make([]int, cols)  // initialize a slice of dx unit8 in each of dy slices
    }
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
            if(i==0 || j==0){
                dp[i][j]=0
                continue
            }
            if(text1[i-1]==text2[j-1]){
                dp[i][j] = dp[i-1][j-1] + 1
            }else{
                dp[i][j] = max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[rows-1][cols-1]
    
    
}
func max (a,b int) int{
    if a>b {
        return a
    }
    return b
}

//Day 6
// func maximalSquare(matrix [][]byte) int {
//     if len(matrix)==0 || matrix==nil{
//         return 0
//     }
//     maxLen:=0
//     rows:=len(matrix)
//     cols:=len(matrix[0])
//     for i:=0;i<rows;i++{
//         for j:=0;j<cols;j++{
//             if matrix[i][j]=='1'{
//                 sqLen:=1
//                 flag:=true
//                 for sqLen+i<rows && sqLen+j<cols && flag==true{
//                     for k:=i;k<=i+sqLen;k++{
//                         if matrix[k][j+sqLen]=='0'{
//                             flag=false
//                             break;
//                         }
//                     }
//                     for k:=j;k<=sqLen+j;k++{
//                         if(matrix[i+sqLen][k]=='0'){
//                             flag=false
//                             break;
//                         }
//                     }
//                     if(flag==true){
//                         sqLen++
//                     }
//                 }
//                 if sqLen>maxLen{
//                     maxLen=sqLen
//                 }
//             }
//         }
//     }
//     return maxLen*maxLen
// }

func maximalSquare(matrix [][]byte) int {
     if len(matrix)==0 || matrix==nil{
        return 0
    }
    rows:=len(matrix)
    cols:=len(matrix[0])
    dp:=make([][]int,rows+1)
    for i:=range dp{
        dp[i]=make([]int,cols+1)
    }
    for i:=range dp{
        dp[i][0]=0
    }
    for j:=range dp[0]{
        dp[0][j]=0
    }
    largest:=0
    for i:=1;i<rows+1;i++{
        for j:=1;j<cols+1;j++{
            if matrix[i-1][j-1]=='1'{
                dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
                if(dp[i][j]>largest){
                    largest=dp[i][j]
                }
            }
        }
    }
    return largest*largest
}
func min(a,b,c int) int{
    if (a<=b &&a<=c){
        return a
    }
    if(b<=a && b<=c){
        return b
    }
    return c
}


//Day 7
type FirstUnique struct {
    index map[int]int
    ll []int
}


func Constructor(nums []int) FirstUnique {
    mb:=FirstUnique{
        index:make(map[int]int),
        ll:[]int{},
    }
    for i:=range(nums){
        mb.Add(nums[i])
    }
    return mb
}


func (this *FirstUnique) ShowFirstUnique() int {
    if(len(this.ll)==0){
        return -1
    }
    for len(this.ll)!=0&& this.index[this.ll[0]]!=1{
       this.ll= this.ll[1:]
    }
    if(len(this.ll)==0){
        return -1
    }
    return this.ll[0]
}


func (this *FirstUnique) Add(value int)  {
    // not yet present
    if this.index[value]==0{
        this.index[value]+=1
        this.ll=append(this.ll,value)
    }else{
        if(len(this.ll)!=0 && this.ll[0]==value){
              this.ll= this.ll[1:]
        }
        this.index[value]++
    }
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */