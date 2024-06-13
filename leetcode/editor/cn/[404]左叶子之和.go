package leetcode

import (
	"testing"
)

/**
给定二叉树的根节点 root ，返回所有左叶子之和。



 示例 1：




输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24


 示例 2:


输入: root = [1]
输出: 0




 提示:


 节点数在 [1, 1000] 范围内
 -1000 <= Node.val <= 1000




 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 708 👎 0

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
func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		//判断是否左叶子节点
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumOfLeftLeaves(t *testing.T) {

}
