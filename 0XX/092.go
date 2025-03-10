package leetcode_ans

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{Val: 0, Next: head}
	head = res
	for i := 0; i < left-1; i++ {
		head = head.Next
	}
	reversed, _ := reverseList(head.Next, right-left+1)
	head.Next = reversed

	return res.Next
}

func reverseList(head *ListNode, count int) (*ListNode, *ListNode) {
	if count == 1 {
		return head, head.Next
	}
	leftHead, rightHead := reverseList(head.Next, count-1)
	head.Next.Next = head
	head.Next = rightHead
	return leftHead, rightHead
}