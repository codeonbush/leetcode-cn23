package leetcode

import (
	"testing"
)

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
//
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 树 二叉搜索树 动态规划 回溯 二叉树 👍 1510 👎 0

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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var build func(l, r int) []*TreeNode
	build = func(l, r int) []*TreeNode {
		res := make([]*TreeNode, 0)
		if l > r {
			res = append(res, nil)
			return res
		}
		for i := l; i <= r; i++ {
			//后序遍历
			leftTree := build(l, i-1)
			rightTree := build(i+1, r)
			//在root上挂上所有可能的左右子树
			for _, left := range leftTree {
				for _, right := range rightTree {
					root := &TreeNode{Val: i}
					root.Left = left
					root.Right = right
					res = append(res, root)
				}
			}
		}
		return res
	}
	return build(1, n)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUniqueBinarySearchTreesIi(t *testing.T) {
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
