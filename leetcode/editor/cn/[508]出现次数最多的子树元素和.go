package leetcode

import (
	"testing"
)

/**
给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。



 示例 1：




输入: root = [5,2,-3]
输出: [2,-3,4]


 示例 2：




输入: root = [5,2,-5]
输出: [2]




 提示:


 节点数在 [1, 10⁴] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 哈希表 二叉树 👍 240 👎 0

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
func findFrequentTreeSum(root *TreeNode) []int {
	var sumToCount map[int]int = make(map[int]int) // 存储子树和及其出现的频率

	// 定义：输入一个节点，返回以该节点为根的二叉树所有节点之和
	var sum func(root *TreeNode) int
	sum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := sum(root.Left)
		rightSum := sum(root.Right)
		res := root.Val + leftSum + rightSum

		// 后序遍历位置，顺手记录子树和对应的频率
		sumToCount[res]++

		return res
	}

	// 遍历二叉树，记录所有子树和及出现频率
	sum(root)

	// 找到最大的出现频率
	maxCount := 0
	for _, count := range sumToCount {
		maxCount = max(maxCount, count)
	}

	// 找到最大出现频率对应的的子树和
	res := []int{}
	for key, value := range sumToCount {
		if value == maxCount {
			res = append(res, key)
		}
	}

	return res
}

// 找最大值函数
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMostFrequentSubtreeSum(t *testing.T) {

}
