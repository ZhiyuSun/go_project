package main


// 543. 合并二叉树
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 2021.05.07 我居然自己做出来了
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val:root1.Val+root2.Val} // 如果这里不加&，return哪里就要加*
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

// 2021.05.07 官方解法，不必开root变量，直接在t1上操作
func mergeTrees1(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func main()  {

}