package leetcode

import (
	"testing"
)

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
//
// Related Topics 栈 树 深度优先搜索 链表 二叉树 👍 1591 👎 0

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
//因为需要在原地构建，遍历方案需要额外变量，最终返回这个额外变量
//递归方案
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//利用定义先将左右子树拉平
	flatten(root.Left)
	flatten(root.Right)
	//将左子树作为右子树
	right := root.Right
	left := root.Left
	root.Right = left
	root.Left = nil
	//用遍历链表的方法，将右子树接在现在的右子树末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
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
