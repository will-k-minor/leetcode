package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// start out at the beginning of each list (x and y)
	// if x <= y, set x.next = y, and next = x.next
	// if y <= x, set y.next = x, and next = y.next
	// if y.next is null, set next to y.next
	// if x.next is null, exit loop

	head := &ListNode{}
	current := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else if list2.Val <= list1.Val {
			current.Next = list2
			list2 = list2.Next
		}
	}

	if list2 != nil {
		current = list2
		list2 = list2.Next
	} else if list1 != nil {
		current = list1
		list1 = list1.Next
	}

	return head.Next
}
