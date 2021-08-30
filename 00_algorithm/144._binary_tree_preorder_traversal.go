package main


// 2021.06.13 好久不练，已经全部不会了，我可能要暂时放弃go了
func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// 2021.06.13 好久不练，已经全部不会了，我可能要暂时放弃go了
func preorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func main() {
	
}
