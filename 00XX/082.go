package leetcode_ans

func deleteDuplicates(head *ListNode) *ListNode {
    temp := &ListNode{0, head}
    prev := temp
    node := head

    for node != nil {
        for node.Next != nil && node.Val == node.Next.Val {
            node = node.Next
        }

        if prev.Next == node {
            prev = prev.Next
        } else {
            prev.Next = node.Next
        }

        node = node.Next
    }

    return temp.Next
}
