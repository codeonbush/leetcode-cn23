package leetcode

import (
	"testing"
)

// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
// 返回这颗 二叉树 。
//
// 示例 1:
//
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
//
// 示例 2:
//
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]
//
// 提示:
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder 和 postorder 都由 不同 的值组成
// postorder 中每一个值都在 inorder 中
// inorder 保证是树的中序遍历
// postorder 保证是树的后序遍历
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1146 👎 0
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	//从后序遍历确定根节点位置：

	//后序遍历数组中，根节点位于最后一个元素，因此根节点的值为 postorder[len(postorder)-1]。
	//在中序遍历数组中，找到根节点的位置，将数组分为左子树和右子树。根节点的位置在中序遍历数组中的索引为 rootIndex。
	//左子树的中序遍历数组：inorder[:rootIndex]
	//右子树的中序遍历数组：inorder[rootIndex+1:]

	//公式：
	//根节点值：rootVal = postorder[len(postorder)-1]
	//左子树后序遍历数组：postorder[:rootIndex]
	//右子树后序遍历数组：postorder[rootIndex:len(postorder)-1]
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 后序遍历的最后一个元素是根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			rootIndex = i
			break
		}
	}

	// 递归构建左子树和右子树
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromInorderAndPostorderTraversal(t *testing.T) {
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
