package leetcode

import (
	"testing"
)

/**
给你一个二叉树的根节点 root ， 检查它是否轴对称。



 示例 1：


输入：root = [1,2,2,3,4,4,3]
输出：true


 示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false




 提示：


 树中节点数目在范围 [1, 1000] 内
 -100 <= Node.val <= 100




 进阶：你可以运用递归和迭代两种方法解决这个问题吗？

 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2725 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

// 定义：判断输入的两棵树是否是镜像对称的
// 1. basecase 根节点相同
// 2. 左右子树镜像对称
func check(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	// 两个根节点需要相同
	if left.Val != right.Val {
		return false
	}
	// 左右子树也需要镜像对称
	return check(left.Right, right.Left) && check(left.Left, right.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSymmetricTree(t *testing.T) {

}
