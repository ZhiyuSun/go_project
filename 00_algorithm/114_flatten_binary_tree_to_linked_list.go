package main


// 114. 二叉树展开为链表

// 2021.05.14 直奔题解，这题不好理解
func flatten(root *TreeNode)  {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func main() {
	
}
