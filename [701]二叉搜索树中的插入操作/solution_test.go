package leetcode

import (
	"testing"
)

//给定二叉搜索树（BST）的根节点
// root 和要插入树中的值
// value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
//
// 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,7,1,3], val = 5
//输出：[4,2,7,1,3,5]
//解释：另一个满足题目要求可以通过的树是：
//
//
//
// 示例 2：
//
//
//输入：root = [40,20,60,10,30,50,70], val = 25
//输出：[40,20,60,10,30,50,70,null,null,25]
//
//
// 示例 3：
//
//
//输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
//输出：[4,2,7,1,3,5]
//
//
//
//
// 提示：
//
//
// 树中的节点数将在
// [0, 10⁴]的范围内。
//
// -10⁸ <= Node.val <= 10⁸
// 所有值
// Node.val 是 独一无二 的。
// -10⁸ <= val <= 10⁸
// 保证 val 在原始BST中不存在。
//
//
// Related Topics 树 二叉搜索树 二叉树 👍 541 👎 0

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空位置插⼊新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// if (root.val == val)
	// BST 中⼀般不会插⼊已存在元素
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInsertIntoABinarySearchTree(t *testing.T) {
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
