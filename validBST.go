package main;

type TreeNodeBST struct {
	Val int
	Left *TreeNodeBST
	Right *TreeNodeBST
}
 
 func isValidBST(root *TreeNodeBST) bool {
    /**
        Recursive

        The tree should be in order

        Go all the way to the left
            start at the lowest key
            we should always be increasing
            if we run into a number that is out of order, its invalid

        Create a stack in order to traverse UPwards

		INORDER TRAVERSAL
    */
	stack := []*TreeNodeBST{}
	inOrder := []int{}

	for root != nil || len(stack) > 0 {
		// dive all the way to the left
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop the last element
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// first value
		inOrder = append(inOrder, root.Val)
		root = root.Right
		if len(inOrder) > 1 && inOrder[len(inOrder)-1] <= inOrder[len(inOrder)-2] {
			return false
		}
	}

	return true;
}