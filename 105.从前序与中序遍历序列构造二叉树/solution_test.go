package leetcode

import (
	"testing"
)

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。
//
//
//
// 示例 1:
//
//
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//
//
//
//
// 提示:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2153 👎 0

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
//二叉树的构造问题一般都是使用「分解问题」的思路：构造整棵树 = 根节点 + 构造左子树 + 构造右子树。

func buildTree(preorder []int, inorder []int) *TreeNode {
	//从先序遍历确定根节点位置：
	//
	//先序遍历数组中，根节点位于第一个元素，因此根节点的值为 preorder[0]。
	//在中序遍历数组中，找到根节点的位置，将数组分为左子树和右子树。根节点的位置在中序遍历数组中的索引为 rootIndex。
	//左子树的中序遍历数组：inorder[:rootIndex]
	//右子树的中序遍历数组：inorder[rootIndex+1:]
	//公式：
	//
	//根节点值：rootVal = preorder[0]
	//左子树先序遍历数组：preorder[1 : 1+rootIndex]
	//右子树先序遍历数组：preorder[1+rootIndex:]

	//如何避免数组越界？
	//在 Go 语言中，切片的长度是指切片中当前元素的个数，而容量是切片的最大容量，cap 值永远大于等于 len 值。
	//因此，如果切片的长度小于或等于切片的容量，那么直接通过类似[0:1]这样的方式来获取切片的子切片是不会导致数组越界的。
	//preorder 和 inorder 数组的长度都不会超过其容量
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//1.找出根节点 preorder[0]
	//2.确定左右子树范围：
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
		}
	}
	root := &TreeNode{Val: preorder[0]}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
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
