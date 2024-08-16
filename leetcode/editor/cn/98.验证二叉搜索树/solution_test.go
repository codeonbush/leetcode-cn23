package leetcode

import (
	"math"
	"testing"
)

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
// 示例 1：
//
// 输入：root = [2,1,3]
// 输出：true
//
// 示例 2：
//
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
// 提示：
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 2065 👎 0
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
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 对于每⼀个节点 root，代码值检查了它的左右孩⼦节点是否符合左⼩右⼤的原则；但
// 是根据 BST 的定义，root 的整个左⼦树都要⼩于 root.val，整个右⼦树都要⼤于 root.val。
// 问题是，对于某⼀个节点 root，他只能管得了⾃⼰的左右⼦节点，需要把 root 的约束传递给左右⼦树
// 限定左⼦树的最⼤值是 root.val，右⼦树的最⼩值是 root.val
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidateBinarySearchTree(t *testing.T) {

}
