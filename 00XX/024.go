package leetcode_ans

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var first *ListNode = head
    var second *ListNode = head.Next
    first.Next = swapPairs(second.Next)
    second.Next = first
    return second
}