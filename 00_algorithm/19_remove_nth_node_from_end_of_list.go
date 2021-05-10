package main


// 19. 删除链表的倒数第 N 个结点

type ListNode struct {
	Val int
	Next *ListNode
}

// 2021.05.08 方法1，计算链表的长度
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next { // 这个地方写(*head).Next也没问题
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 2021.05.08 方法2，利用栈
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 2021.05.08 方法3，双指针。这解法也太神奇了
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main()  {

}