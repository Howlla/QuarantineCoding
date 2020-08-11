//Day 1
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    _,triangle:=calculate(root);
    return triangle;
}

func calculate(root *TreeNode) (int,int){
    if root==nil{
        return 0,-99999;
    }
    branchLeft,subTreeLeft:= calculate(root.Left)
    branchRight,subTreeRight:=calculate(root.Right)
    maxSubBranch:=max(branchLeft,branchRight)
    val:=root.Val
    maxAsBranch:=max(val,maxSubBranch+val)
    maxTotal:=max(maxAsBranch,branchLeft+val+branchRight)
    ans:=max(subTreeLeft,subTreeRight,maxTotal)
    return maxAsBranch,ans
}

func max(first int, vals ...int) int{
	for _, val:=range vals{
		if val>first{
			first=val
		}
	}
	return first
}

//Day 2
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidSequence(root *TreeNode, arr []int) bool {
    return dfs(root,arr,len(arr),0);
}

func dfs(root *TreeNode, arr []int, n,i int) bool{
    //When do we fail?
    if(arr[i]!=root.Val){
        return false
    }
    if(i==n-1){
        return root.Left==nil && root.Right==nil
    }
    return (root.Left!=nil && dfs(root.Left,arr,n,i+1)) || (root.Right!=nil && dfs(root.Right,arr,n,i+1))
}








