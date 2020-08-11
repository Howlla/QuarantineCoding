//Day 1
func productExceptSelf(nums []int) []int {
    n:=len(nums)
    left:=make([]int,n)
    right:=make([]int,n)
    left[0]=1
    for i:=1;i<n;i++{
        left[i]=left[i-1]*nums[i-1]
    }
    right[n-1]=1
    for i:=n-2;i>=0;i--{
        right[i]=right[i+1]*nums[i+1]
    }
    output:=make([]int,n)
    for i:=0;i<n;i++ {
        output[i] = left[i]*right[i]
    }
    return output
}

//Day 2

func checkValidString(s string) bool {
    stack1:=new(Stack)
    stack2:=new(Stack)
    for i:=0;i<len(s);i++{
        if(s[i]=='('){
            stack1.Push(i)
        }else if(s[i]=='*'){
            stack2.Push(i)
        }else{
            if!stack1.isEmpty(){
                stack1.Pop()
            }else if !stack2.isEmpty(){
                stack2.Pop()
            }else{
                return false
            }
        }
    }
    if(stack1.isEmpty()){
        return true
    }
    for (!stack1.isEmpty()){
        x:=stack1.Pop()
        if(stack2.isEmpty()){
           return false 
        }else{
            y:=stack2.Pop()
            if(y<x){
                return false
            }
        }
    }
    return true
}

type Stack struct{
    slice []int
}

func (s *Stack) Push(x int){
    s.slice=append(s.slice,x)
}
func (s *Stack) Peek()int{
    if(len(s.slice)==0){
        return -1
    }
    return s.slice[len(s.slice)-1]
}
func (s *Stack) isEmpty()bool{
    if(len(s.slice)==0){
        return true
    }
    return false
}
func (s *Stack) Pop() int{
    if(len(s.slice)==0){
        return -1
    }
    x:=s.Peek()
    s.slice = s.slice[0:len(s.slice)-1]
    return x
}

//Day 3

func numIslands(grid [][]byte) int {
    count:=0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if(grid[i][j]=='1'){
                count++;
                dfs(grid,i,j,len(grid),len(grid[0]))
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i,j,n,m int){
    if(i>=n || i<0 || j>=m||j<0){
        return
    }
    if(grid[i][j]=='1'){
        grid[i][j]='0'
        dfs(grid,i+1,j,n,m)
        dfs(grid,i,j+1,n,m)
        dfs(grid,i-1,j,n,m)
        dfs(grid,i,j-1,n,m)
    }else{
        return
    }
}

//Day 4
func minPathSum(grid [][]int) int {
    n:= len(grid)
    m:= len(grid[0])
    dp:=make([][]int,n)
    for i:= range dp{
        dp[i]=make([]int,m)
    }
    dp[0][0]=grid[0][0]
    for i:=1;i<n;i++{
        dp[i][0]=dp[i-1][0]+grid[i][0]
    }
    for j:=1;j<m;j++{
        dp[0][j]=dp[0][j-1]+grid[0][j]
    }
    for i:=1;i<n;i++{
        for j:=1;j<m;j++{

                dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
            
        }
    }
    fmt.Println(dp)
    return dp[n-1][m-1]
}

func min (a,b int)int{
    if (a<b){
        return a
    }
    return b
}

//Day 5
class Solution {
    public int search(int[] A, int target) {
         int n=A.length;
         int lo=0,hi=n-1;
        // find the index of the smallest value using binary search.
        // Loop will terminate since mid < hi, and lo or hi will shrink by at least 1.
        // Proof by contradiction that mid < hi: if mid==hi, then lo==hi and loop would have been terminated.
        while(lo<hi){
            int mid=(lo+hi)/2;
            if(A[mid]>A[hi]) lo=mid+1;
            else hi=mid;
        }
        // lo==hi is the index of the smallest value and also the number of places rotated.
        int rot=lo;
        lo=0;hi=n-1;
        // The usual binary search and accounting for rotation.
        while(lo<=hi){
            int mid=(lo+hi)/2;
            int realmid=(mid+rot)%n;
            if(A[realmid]==target)return realmid;
            if(A[realmid]<target)lo=mid+1;
            else hi=mid-1;
        }
        return -1;
    }
}

//Day 6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    if(len(preorder)==0){
        return nil
    }
    n:= new(TreeNode)  
    n.Val=preorder[0]
    i:=0
    for i<len(preorder)&&preorder[i]<=preorder[0]{
        i++
    }
    n.Left=bstFromPreorder(preorder[1:i])
    n.Right=bstFromPreorder(preorder[i:])
    return n
}

//Day 7
/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    dims:=binaryMatrix.Dimensions();
    n:=dims[0]
    m:=dims[1]
    ans:=-1
    i:=n-1
    j:=m-1
    for i>=0 && j>=0 {
        if(binaryMatrix.Get(i,j)==1){
            ans=j
            j--
        }else{
            i--
        }
    }
    return ans
}

func min(a,b int) int{
    if a<b{
        return a
    }
    return b
}