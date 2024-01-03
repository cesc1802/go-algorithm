package middle_linked_list

import "algorithm/structure"

func LookingForMiddleNode(head *structure.LinkedListNode) *structure.LinkedListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
