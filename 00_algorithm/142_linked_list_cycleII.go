package main


// 142. 环形链表 II

// 2021.05.22 模仿python的方法写的
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}

// 这个数学公式的推导太牛逼了，一看就懂
//slow * 2 = fast;
//slow = a + b;
//fast = a + b + c + b = a + 2*b + c;
//(a + b)*2 = a + 2*b + c;
//a = c;
//口头解释就是
//快针走的是慢针的两倍。
//慢针走过的路，快针走过一遍。
//快针走过的剩余路程，也就是和慢针走过的全部路程相等。(a+b = c+b)
//刨去快针追赶慢针的半圈(b)，剩余路程即为所求入环距离(a=c)


func main() {
	
}
