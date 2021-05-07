package main

import "fmt"

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func printListNode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next //移动指针
	}
}

func main()  {
	l1 := &ListNode{Val: 2}
	head1 := l1
	l1.Next = &ListNode{Val: 4}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	head2 := l2
	l2.Next = &ListNode{Val: 6}
	l2 = l2.Next
	l2.Next = &ListNode{Val: 4}

	head := addTwoNumbers(head1, head2)
	printListNode(head)
}