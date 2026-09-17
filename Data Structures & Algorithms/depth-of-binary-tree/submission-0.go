/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
    leftMax := depthCount(root.Left, 1)
	rightMax := depthCount(root.Right, 1)
	return max(leftMax, rightMax)
}

func depthCount(root *TreeNode, n int) int {
	if root == nil {
		return n
	}
	n++
	leftMax := depthCount(root.Left, n)
	rightMax := depthCount(root.Right, n)
	return max(leftMax, rightMax)
}