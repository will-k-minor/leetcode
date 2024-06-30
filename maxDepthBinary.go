package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(leaf *TreeNode, depth int, maxDepth *int)

	dfs = func(leaf *TreeNode, depth int, maxDepth *int) {
		// if the current leaf is null, we have hit the end
		// set the max to the maximum depth so far
		if leaf == nil {
			if depth > *maxDepth {
				*maxDepth = depth
			}
			return
		}
		dfs(leaf.Left, depth+1, maxDepth)
		dfs(leaf.Right, depth+1, maxDepth)
	}

	if root == nil {
		return 0
	}

	max := 0
	dfs(root, 0, &max)
	return max
}
