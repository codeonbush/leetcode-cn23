package leetcode

import (
	"testing"
)

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
// 示例 1：
//
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//
// 示例 2：
//
// 输入：root = [2,1,3]
// 输出：[2,3,1]
//
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目范围在 [0, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1735 👎 0
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

// 递归解法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	var temp *TreeNode
	temp = root.Left
	root.Left = root.Right
	root.Right = temp
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInvertBinaryTree(t *testing.T) {
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
