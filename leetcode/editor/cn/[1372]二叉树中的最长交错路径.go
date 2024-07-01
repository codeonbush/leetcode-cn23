package leetcode

import (
	"testing"
)

/**
给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：


 选择二叉树中 任意 节点和一个方向（左或者右）。
 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
 改变前进方向：左变右或者右变左。
 重复第二步和第三步，直到你在树中无法继续移动。


 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。

 请你返回给定树中最长 交错路径 的长度。



 示例 1：



 输入：root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
输出：3
解释：蓝色节点为树中最长交错路径（右 -> 左 -> 右）。


 示例 2：



 输入：root = [1,1,1,null,1,null,null,1,1,null,1]
输出：4
解释：蓝色节点为树中最长交错路径（左 -> 右 -> 左 -> 右）。


 示例 3：

 输入：root = [1]
输出：0




 提示：


 每棵树最多有 50000 个节点。
 每个节点的值在 [1, 100] 之间。


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 171 👎 0

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
//如果一个从上到下的交错路径的开头是从右到左的，称之为「左交错路径」，反之成为「右交错路径」。
//
//这样的话，一个节点 x 能够产生的交错路径就能分解到左右子树：
//
//1、x 的左子树的「右交错路径」+ 节点 x = x 的「左交错路径」
//
//2、x 的右子树的「左交错路径」+ 节点 x = x 的「右交错路径」
//
//比较 x 的左右交错路径，即可算出以 x 开头的最长交错路径。
//
//对二叉树上的所有节点计算交错路径长度，即可得到最长的交错路径长度。
//
//所以我们定义一个 getPathLen 函数计算并返回一个节点的左右交错路径长度，然后在后序位置上更新全局最大值。

// 关键是把交错路径分成左右两个值来考虑
func longestZigZag(root *TreeNode) int {
	var res int

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getPathLen func(root *TreeNode) []int
	getPathLen = func(root *TreeNode) []int {
		if root == nil {
			return []int{-1, -1}
		}
		left := getPathLen(root.Left)
		right := getPathLen(root.Right)
		// 后序位置，根据左右子树的交错路径长度推算根节点的交错路径长度
		rootPathLen1 := left[1] + 1
		rootPathLen2 := right[0] + 1
		// 更新全局最大值
		res = max(res, max(rootPathLen1, rootPathLen2))

		return []int{rootPathLen1, rootPathLen2}
	}

	getPathLen(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestZigzagPathInABinaryTree(t *testing.T) {

}
