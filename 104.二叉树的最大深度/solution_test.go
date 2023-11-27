package leetcode

import (
	"testing"
)

// 给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例： 给定二叉树 [3,9,20,null,null,15,7]，
//
//	   3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// 返回它的最大深度 3 。
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1631 👎 0
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//1. 遍历解法
// 记录最大深度

func maxDepth(root *TreeNode) int {
	var res int
	// 记录遍历到的节点的深度
	var depth int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 前序位置
		depth++
		if root.Left == nil && root.Right == nil {
			// 到达叶子节点，更新最大深度
			res = max(res, depth)
		}
		traverse(root.Left)
		traverse(root.Right)
		//后序位置
		depth--
	}
	traverse(root)
	return res
}

// 2. 递归解法
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	// 构建二叉树 [1,null,2]
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	maxdep := maxDepth(root)
	print(maxdep)
}
