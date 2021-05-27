package main

// 141. 环形链表

// 2021.05.22 参考python解法自己写的
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	a, b := head, head.Next
	for a != nil && b != nil && b.Next != nil {
		if a == b {
			return true
		}
		a = a.Next
		b = b.Next.Next
	}
	return false
}


// 官方解法1，哈希表
func hasCycle1(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 我不太清楚为甚么是struct{}{}，如果换成int{}可以吗
func hasCycle2(head *ListNode) bool {
	seen := map[*ListNode]int{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = 1
		head = head.Next
	}
	return false
}

func hasCycle4(head *ListNode) bool {
	seen := map[*ListNode]bool{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = true
		head = head.Next
	}
	return false
}



// 官方解法2，快慢指针
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	
}
