package leetcode_ans

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	l := listLen(head)

    if l == 0 {
		return head
	}

	k %= l

	if k == 0 {
		return head
	}

	prev := l - k

	nodePrev, nodeLast := getNodes(head, prev)

	nodePos := nodePrev.Next

	nodePrev.Next = nil
	nodeLast.Next = head

	return nodePos
}

func getNodes(head *ListNode, prev int) (*ListNode, *ListNode) {
	index := 1

	var nodePrev *ListNode

	for head.Next != nil {
		if index == prev {
			nodePrev = head
		}

		index++
		head = head.Next
	}

	return nodePrev, head
}

func listLen(node *ListNode) int {
	l := 0

	for node != nil {
		node = node.Next
		l++
	}

	return l
}