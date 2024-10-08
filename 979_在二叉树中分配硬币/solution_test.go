package leetcode

import (
	"testing"
)

/**
给你一个有 n 个结点的二叉树的根结点 root ，其中树中每个结点 node 都对应有 node.val 枚硬币。整棵树上一共有 n 枚硬币。

 在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。移动可以是从父结点到子结点，或者从子结点移动到父结点。

 返回使每个结点上 只有 一枚硬币所需的 最少 移动次数。



 示例 1：


输入：root = [3,0,0]
输出：2
解释：一枚硬币从根结点移动到左子结点，一枚硬币从根结点移动到右子结点。


 示例 2：


输入：root = [0,3,0]
输出：3
解释：将两枚硬币从根结点的左子结点移动到根结点（两次移动）。然后，将一枚硬币从根结点移动到右子结点。




 提示：


 树中节点的数目为 n
 1 <= n <= 100
 0 <= Node.val <= n
 所有 Node.val 的值之和是 n


 Related Topics 树 深度优先搜索 二叉树 👍 532 👎 0

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
func distributeCoins(root *TreeNode) int {
	res := 0
	var getRest func(*TreeNode) int
	getRest = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getRest(root.Left)
		right := getRest(root.Right)
		// 让当前这棵树平衡所需的移动次数
		res += abs(left) + abs(right) + (root.Val - 1)
		// 实现函数的定义
		return left + right + (root.Val - 1)
	}
	getRest(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCoinsInBinaryTree(t *testing.T) {

}
