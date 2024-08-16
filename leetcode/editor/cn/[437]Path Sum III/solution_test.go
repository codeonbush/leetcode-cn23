package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree and an integer targetSum, return the number of
paths where the sum of the values along the path equals targetSum.

 The path does not need to start or end at the root or a leaf, but it must go
downwards (i.e., traveling only from parent nodes to child nodes).


 Example 1:


Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.


 Example 2:


Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3



 Constraints:


 The number of nodes in the tree is in the range [0, 1000].
 -10⁹ <= Node.val <= 10⁹
 -1000 <= targetSum <= 1000


 Related Topics 树 深度优先搜索 二叉树 👍 1861 👎 0

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
// 记录前缀和
// 定义：从二叉树的根节点开始，路径和为 pathSum 的路径有 preSumCount[pathSum] 个
func pathSum(root *TreeNode, targetSum int) int {
	//前缀和
	preSumCount := make(map[int]int)
	//路径和
	pathSum, res := 0, 0
	preSumCount[0] = 1
	traverseT(root, targetSum, &pathSum, &res, preSumCount)
	return res
}

func traverseT(root *TreeNode, targetSum int, pathSum, res *int, preSumCount map[int]int) {
	if root == nil {
		return
	}
	// 前序遍历位置
	*pathSum += root.Val
	// 从二叉树的根节点开始，路径和为 pathSum - targetSum 的路径条数
	// 就是路径和为 targetSum 的路径条数
	*res += preSumCount[*pathSum-targetSum]
	// 记录从二叉树的根节点开始，路径和为 pathSum 的路径条数
	preSumCount[*pathSum]++
	traverseT(root.Left, targetSum, pathSum, res, preSumCount)
	traverseT(root.Right, targetSum, pathSum, res, preSumCount)
	// 后序遍历位置
	preSumCount[*pathSum]--
	*pathSum -= root.Val
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSumIii(t *testing.T) {

}
