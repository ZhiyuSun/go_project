package main

// 543. 二叉树的直径
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 学习go语言全局变量的使用
var max = 0
func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	if root == nil{
		return 0
	}
	deep(root)
	return max
}

func deep(root *TreeNode)int{
	if root == nil{
		return 0
	}

	left :=deep(root.Left)
	right := deep(root.Right)
	if max < left+right {
		max = left+right
	}
	return maxV(left,right)+1

}

func maxV(a,b int)int  {
	if a>=b{
		return a
	}else {
		return b
	}
}

func main()  {
	
}