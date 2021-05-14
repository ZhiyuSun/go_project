package main

// 102. 二叉树的层序遍历

// 2021.05.14 我的最开始的写法，感觉有些问题
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]int, 0)
		l := len(queue)
		for i:=0; i<l; i++ {  // 我原来写的是len(queue),结果一直在动态变化
			tmp = append(tmp, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, append([]int{}, tmp...))
	}
	return res
}

// 官方解法
func levelOrder1(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}


func main()  {

}