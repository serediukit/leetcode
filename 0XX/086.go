package leetcode_ans

func partition(head *ListNode, x int) *ListNode {
    if head == nil || x <= -100 || x > 100 {
        return head
    }

    less := &ListNode{Val:0}
    greater := &ListNode{Val:0}

    lessStart := less
    greaterStart:= greater

    for head != nil {
        if head.Val >= x {
            greater.Next = head
            greater = greater.Next
        } else {
            less.Next = head
            less = less.Next
        }
        head = head.Next
    }

    greater.Next = nil
    less.Next = greaterStart.Next

    return lessStart.Next
}