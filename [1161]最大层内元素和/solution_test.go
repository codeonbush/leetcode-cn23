package leetcode

import (
	"math"
	"testing"
)

/**
给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。

 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。



 示例 1：




输入：root = [1,7,0,7,-8,null,null]
输出：2
解释：
第 1 层各元素之和为 1，
第 2 层各元素之和为 7 + 0 = 7，
第 3 层各元素之和为 7 + -8 = -1，
所以我们返回第 2 层的层号，它的层内元素之和最大。


 示例 2：


输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
输出：2




 提示：


 树中的节点数在
 [1, 10⁴]范围内

 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 137 👎 0

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
//key1 记录层数，元素和最大的层
//这道题实际解法只返回一层即可
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	// 记录 BFS 走到的层数
	depth := 1
	// 记录元素和最大的那一行和最大元素和
	res, maxSum := 0, int(math.MinInt32)

	for len(q) != 0 {
		sz := len(q)
		levelSum := 0
		// 遍历这一层
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			levelSum += cur.Val

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		if levelSum > maxSum {
			// 更新最大元素和
			res = depth
			maxSum = levelSum
		}
		depth++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumLevelSumOfABinaryTree(t *testing.T) {

}
