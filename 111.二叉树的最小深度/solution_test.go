package leetcode

import (
	"testing"
)

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
// 示例 2：
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
// 提示：
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1132 👎 0
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
func minDepth(root *TreeNode) int {
	minDep := 0
	depth := 0
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++
		if node.Left == nil && node.Right == nil {
			if depth > 0 && minDep > 0 {
				minDep = min(minDep, depth)
			} else {
				minDep = depth
			}
		}
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}
	traverse(root)
	return minDep
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumDepthOfBinaryTree(t *testing.T) {
	//cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d expected %d, but %d got",
	//				c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}
}
