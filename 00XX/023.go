package leetcode_ans

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    for i := 1; i < len(lists); i *= 2 {
        for j := 0; j < len(lists) - i; j += i*2 {
            lists[j] = mergeTwoLists(lists[j], lists[j+i])
        }
    }
    return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    } else if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2);
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next);
        return list2
    }
}