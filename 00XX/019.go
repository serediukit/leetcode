package leetcode_ans

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    zeroNode := &ListNode{Next: head}
	currentNode := head
	var previousNode *ListNode = zeroNode
	for currentNode != nil {
		if numberFromEnd(currentNode) == n {
			previousNode.Next = currentNode.Next
			break
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return zeroNode.Next
}

func numberFromEnd(node *ListNode) int {
	if node.Next != nil {
		return 1 + numberFromEnd(node.Next)
	}
	return 1
}