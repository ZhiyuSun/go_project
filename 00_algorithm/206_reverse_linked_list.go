package main

// 206. 反转链表

// 2021.05.10 自己做出来了
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 官方解法1，迭代法
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main()  {

}
