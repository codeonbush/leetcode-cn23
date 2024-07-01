package leetcode

import (
	"testing"
)

/**
给定一个非空二叉树的根节点
 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10⁻⁵ 以内的答案可以被接受。



 示例 1：




输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。


 示例 2:




输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]




 提示：





 树中节点数量在 [1, 10⁴] 范围内
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 493 👎 0

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

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		// 记录当前层所有节点之和
		sum := 0.0
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			sum += float64(cur.Val)
		}
		// 记录当前行的平均值
		res = append(res, sum/float64(size))
	}

	return res
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		var levelAvg float64
		var levelTotal float64
		for i := 0; i < size; i++ {
			cur := q[0]
			levelTotal += float64(cur.Val)
			levelAvg = levelTotal / float64(size)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			q = q[1:]
		}
		res = append(res, levelAvg)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAverageOfLevelsInBinaryTree(t *testing.T) {

}
