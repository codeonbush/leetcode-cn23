package leetcode

import (
	"math"
	"testing"
)

/**
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过
根节点。

 路径和 是路径中各节点值的总和。

 给你一个二叉树的根节点 root ，返回其 最大路径和 。



 示例 1：


输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

 示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42




 提示：


 树中节点数目范围是 [1, 3 * 10⁴]
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 2230 👎 0

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
//分解问题递归和迭代递归的区别是分解问题函数有返回值
//最大路径和
//最大单边路径和
//最大单边路径和与计算最大深度相似，区别是加上val，最大深度是+1
func maxPathSum(root *TreeNode) int {
	var res = math.MinInt32

	// 定义：计算从根节点 root 为起点的最大单边路径和
	var oneSideMax func(*TreeNode) int
	oneSideMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		//最大路径和至少为0
		leftMaxSum := int(math.Max(float64(0), float64(oneSideMax(root.Left))))
		rightMaxSum := int(math.Max(float64(0), float64(oneSideMax(root.Right))))

		// 后序遍历位置，顺便更新最大路径和
		pathMaxSum := root.Val + leftMaxSum + rightMaxSum
		res = int(math.Max(float64(res), float64(pathMaxSum)))

		// 实现函数定义，左右子树的最大单边路径和加上根节点的值
		// 就是从根节点 root 为起点的最大单边路径和
		return int(math.Max(float64(leftMaxSum), float64(rightMaxSum))) + root.Val
	}

	if root == nil {
		return 0
	}

	// 计算单边路径和时顺便计算最大路径和
	oneSideMax(root)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeMaximumPathSum(t *testing.T) {

}
