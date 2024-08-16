package leetcode

import (
	"fmt"
	"testing"
)

// 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵
// 树的后序遍历，重构并返回二叉树。
//
// 如果存在多个答案，您可以返回其中 任何 一个。
//
// 示例 1：
//
// 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]
//
// 示例 2:
//
// 输入: preorder = [1], postorder = [1]
// 输出: [1]
//
// 提示：
//
// 1 <= preorder.length <= 30
// 1 <= preorder[i] <= preorder.length
// preorder 中所有值都 不同
// postorder.length == preorder.length
// 1 <= postorder[i] <= postorder.length
// postorder 中所有值都 不同
// 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 330 👎 0
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{postorder[0], nil, nil}
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for i = 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			break
		}
	}
	//将两个序列联系起来的变量是子树序列的长度
	//root.Left = constructFromPrePost(preorder[1:preIndex], postorder[:postIndex])
	//root.Right = constructFromPrePost(preorder[preIndex:], postorder[postIndex+1:])
	root.Left = constructFromPrePost(preorder[1:i+2], postorder[:i+1])
	root.Right = constructFromPrePost(preorder[i+2:], postorder[i+1:len(postorder)-1])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromPreorderAndPostorderTraversal(t *testing.T) {
	preorder := []int{2, 1}
	postorder := []int{1, 2}

	root := constructFromPrePost(preorder, postorder)

	fmt.Sprintln(root)

}
