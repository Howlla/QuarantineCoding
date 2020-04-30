//Day 1
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow:= head;
    counter:=0
    for  slow.Next!=nil {
        counter++
        slow=slow.Next
    }
    fast:=head;
    if(counter%2==0){
    counter=counter/2
    }else {
        counter=(counter/2) +1
    }
    
    for ;counter>0;counter-- {
        fast=fast.Next
    }
    return fast
}

//Day 2
func backspaceCompare(S string, T string) bool {
    buf1 := make([]byte, 0)
    for i := range S {
        if S[i] == '#' {
            if len(buf1) > 0 {
                buf1 = buf1[:len(buf1)-1]
            }
            continue
        }   
        buf1 = append(buf1, S[i])
    }
    
    buf2 := make([]byte, 0)
    for i := range T {
        if T[i] == '#' {
            if len(buf2) > 0 {
                buf2 = buf2[:len(buf2)-1]
            }
            continue
        }
        buf2 = append(buf2, T[i])
    }
    
    return string(buf1) == string(buf2)
}

//Day 3
type MinStack struct {
    stack1,stack2 []int;
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
		stack1: make([]int, 0),
		stack2:   make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
    this.stack1=append(this.stack1,x)
    if(len(this.stack2)==0||x<this.stack2[len(this.stack2)-1]){
        this.stack2=append(this.stack2,x)
        
    }else{
        this.stack2=append(this.stack2,this.stack2[len(this.stack2)-1])
    }    
}


func (this *MinStack) Pop()  {
    if len(this.stack1) > 0 {
    this.stack1 = this.stack1[:len(this.stack1)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    }
}


func (this *MinStack) Top() int {
    return this.stack1[len(this.stack1)-1]
}


func (this *MinStack) GetMin() int {
    return this.stack2[len(this.stack2)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

 //Day 4

 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if(root==nil){
        return 0
    }
    ans:=0
    calculate(root,&ans)
    return ans-1
}

func calculate(root *TreeNode, ans* int) int {
    if(root==nil){
        return 0
    }
    leftH:=calculate(root.Left, ans)
    rightH:=calculate(root.Right,ans)
    *ans=max(leftH+rightH+1,*ans)
    return (1+max(leftH,rightH))
}

func max(a,b int)int{
    if(a>b){
        return a
    }
    return b;
}

//Day 5
import "sort"
func lastStoneWeight(stones []int) int {
    sort.Ints(stones);
    var a,b int;
    for {
        if(len(stones)==0){
            return 0
        }else if (len(stones)==1){
            return stones[0]
        }
        a=stones[len(stones)-1]
        b=stones[len(stones)-2]
        if(a==b){
            stones=stones[:len(stones)-2]
        }else{
            stones = stones[:len(stones)-1]
            stones[len(stones)-1] = abs(a,b)   
        }
        sort.Ints(stones);
    }
}

func abs(a,b int) int{
    if(a>b){
        return a-b
    }else{
        return b-a
    }
}

//Day 6
func findMaxLength(nums []int) int {
    hm:=make(map[int]int)
    count:=0
    maxLen:=0
    for i:=range nums{
        if (nums[i]==0){
            count--
        }else{
            count++
        }
        if (count==0){
            maxLen=max(maxLen,i+1)
        }
        if val,ok:=hm[count];ok{
            maxLen=max(maxLen,i-val)
        }else{
            hm[count]=i
        }
    }
    return maxLen
}
func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}

//Day 7
import "strings"

func stringShift(s string, shift [][]int) string {
    finalShift:=0
    for i:=0;i<len(shift);i++{
        if(shift[i][0]==0){
            finalShift-=shift[i][1]
        }else{
            finalShift+=shift[i][1]
        }
    }
    fmt.Println(finalShift)
    finalShift=abs(finalShift,len(s))
    fmt.Println(finalShift)
    k := strings.Split(s,"")
    for i:=0;i<finalShift;i++{
        last:=k[len(k)-1]
        k = k[:len(k)-1]
        k=append([]string{last},k...)
    }
    return strings.Join(k,"")
}
func abs (a,l int) int{
    if(a>=0){
        return (a) %l
    }else{
        return (l+a) %l
    }
}