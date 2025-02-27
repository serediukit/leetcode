package leetcode_ans

func deleteDuplicates(head *ListNode) *ListNode {
    node := head

    for node != nil {
        for node.Next != nil && node.Val == node.Next.Val {
            node.Next = node.Next.Next
        }

        node = node.Next
    }

    return head
}