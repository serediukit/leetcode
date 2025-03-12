package leetcode_ans

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    parts, l := getParts(head, k)
    var reserved = reverseParts(parts, l % k == 0)
    return mergeParts(reserved)
}

func getParts(head *ListNode, k int) ([]*ListNode, int) {
	var lists []*ListNode
	index := 0
	for head != nil {
		nextNode := head.Next
		if index%k == 0 {
			lists = append(lists, head)
		}
		if (index+1)%k == 0 {
			head.Next = nil
		}
		index++
		head = nextNode
	}
	return lists, index
}

func reverseParts(parts []*ListNode, reverseLast bool) []*ListNode {
	var reversed []*ListNode
	var x int
	if !reverseLast {
		x = 1
	}
	for i := 0; i < len(parts)-x; i++ {
		currentNode := parts[i]
		nextNode := parts[i].Next
		currentNode.Next = nil
		for nextNode != nil {
			tempNode := nextNode.Next
			nextNode.Next = currentNode
			currentNode = nextNode
			nextNode = tempNode
		}
		reversed = append(reversed, currentNode)

	}
	if !reverseLast {
		reversed = append(reversed, parts[len(parts)-1])
	}
	return reversed
}

func mergeParts(lists []*ListNode) *ListNode {
	for i := len(lists) - 2; i >= 0; i-- {
		lastNode := lists[i]
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		fmt.Println("LastNode:", lastNode)
		fmt.Println("NextNode:", lists[i+1])
		lastNode.Next = lists[i+1]
	}
	return lists[0]
}