package main

// 538. 把二叉搜索树转换为累加树

// 2021.05.28 模仿中序遍历的解法
func convertBST(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	s := 0
	dumpy := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		root = stack[len(stack)-1]
		s += root.Val
		root.Val = s
		stack = stack[:len(stack)-1]
		root = root.Left
	}
	return dumpy
}


// 2021.05.28 反序中序遍历
func convertBST1(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}


func main() {
	
}
