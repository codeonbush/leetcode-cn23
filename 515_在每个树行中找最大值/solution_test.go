package leetcode

import (
	"math"
	"testing"
)

/**
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。



 示例1：




输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]


 示例2：


输入: root = [1,2,3]
输出: [1,3]




 提示：


 二叉树的节点个数的范围是 [0,10⁴]

 -2³¹ <= Node.val <= 2³¹ - 1




 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 367 👎 0

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
//首先，这题肯定可以用 BFS 算法解决，for 循环里面判断最大值就行了
//遍历的过程中记录对应深度的最大节点值即可
//层序遍历实际是两层循环，第一层从上到下，第二层从左到右
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	// while 循环控制从上向下一层层遍历
	for len(q) > 0 {
		sz := len(q)
		// 记录这一层的最大值
		levelMax := math.MinInt32
		// for 循环控制每一层从左向右遍历
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			levelMax = max(levelMax, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, levelMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestValues_DFS(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 一定要用 array 存储，因为要用索引随机访问
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, root.Val)
		} else {
			// 记录当前行的最大值
			res[depth] = max(res[depth], root.Val)
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindLargestValueInEachTreeRow(t *testing.T) {

}
