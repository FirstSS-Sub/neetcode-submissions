/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	now := ans

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			now.Next = list1
			list1 = list1.Next
		} else {
			now.Next = list2
			list2 = list2.Next
		}
		now = now.Next
	}

	if list1 != nil {
		now.Next = list1
	} else {
		now.Next = list2
	}

	return ans.Next
}
