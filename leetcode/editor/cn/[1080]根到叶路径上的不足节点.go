package leetcode

import (
	"testing"
)

/**
给你二叉树的根节点 root 和一个整数 limit ，请你同时删除树中所有 不足节点 ，并返回最终二叉树的根节点。

 假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为 不足节点 ，需要被删除。

 叶子节点，就是没有子节点的节点。



 示例 1：


输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
输出：[1,2,3,4,null,null,7,8,9,null,14]


 示例 2：


输入：root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
输出：[5,4,8,11,null,17,4,7,null,null,null,5]


 示例 3：


输入：root = [1,2,-3,-5,null,4,null], limit = -1
输出：[1,null,-3,4]




 提示：


 树中节点数目在范围 [1, 5000] 内
 -10⁵ <= Node.val <= 10⁵
 -10⁹ <= limit <= 10⁹




 Related Topics 树 深度优先搜索 二叉树 👍 178 👎 0

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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	// 定义：输入一个节点 root，和约束 limit，
	// 删除以 root 为根的二叉树中的「不足节点」，返回删除完成后的二叉树根节点
	if root == nil {
		return nil
	}
	// 前序位置，接收父节点传递的 limit 约束决定叶子结点是否需要被删除
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			// 对于叶子节点，如果低于 limit 说明需要被删除
			return nil
		}
		return root
	}
	// 先对左右子树进行删除，接收返回值
	left := sufficientSubset(root.Left, limit-root.Val)
	right := sufficientSubset(root.Right, limit-root.Val)
	// 后序位置，根据子树的删除情况决定自己是否需要被删除
	if left == nil && right == nil {
		// 如果左右子树不满足 limit - root.val 的约束，那么就存在经过 root
		// 节点的路径不满足约束，也就说明 root 节点是「不足节点」，需要被删掉
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInsufficientNodesInRootToLeafPaths(t *testing.T) {

}
