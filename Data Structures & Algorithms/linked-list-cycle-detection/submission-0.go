/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
    set := make(map[*ListNode]struct{}, 0)
	_, ok := set[head]
	for head.Next != nil && !ok {
		set[head] = struct{}{}
		head = head.Next
		_, ok = set[head]
	}
	return ok
}
