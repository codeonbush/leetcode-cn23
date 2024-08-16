package leetcode

import (
	"math"
	"testing"
)

/**
给定一个二叉树，判断它是否是 平衡二叉树



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true


 示例 2：


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false


 示例 3：


输入：root = []
输出：true




 提示：


 树中的节点数在范围 [0, 5000] 内
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 👍 1519 👎 0

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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDeptht(root.Left)
	right := maxDeptht(root.Right)
	return math.Abs(float64(left-right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDeptht(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDeptht(root.Left)), float64(maxDeptht(root.Right))) + 1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBalancedBinaryTree(t *testing.T) {

}
