package main

// 104. 二叉树的最大深度

// 2021.05.14 慢慢开始学会写树的题目
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxValue(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 官方解法，广度优先的方法
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue) // 这里也是一个道理，把长度先取出来，不要在for动态的变化长度
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}


func main()  {

}
