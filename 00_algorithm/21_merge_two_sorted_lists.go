package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l1, l2 = l2, l1
		}
		l1.Next = mergeTwoLists(l1.Next, l2)
	}
	if l1 != nil {
		return l1
	} else {
		return l2
	}
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		}else{
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}

func main()  {
	fmt.Println("hello world")
}