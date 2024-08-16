package leetcode

import (
	"testing"
)

/**
给定一个二叉树，我们在树的节点上安装摄像头。

 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

 计算监控树的所有节点所需的最小摄像头数量。



 示例 1：



 输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。


 示例 2：



 输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。


 提示：


 给定树的节点数的范围是 [1, 1000]。
 每个节点的值都是 0。


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 724 👎 0

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
func minCameraCover(root *TreeNode) int {
	var res int
	setCamera(root, false, &res)
	return res
}

// 定义：输入以 root 为根的二叉树，以最优策略在这棵二叉树上放置摄像头，
// 然后返回 root 节点的情况：
// 返回 -1 代表 root 为空，返回 0 代表 root 未被 cover，
// 返回 1 代表 root 已经被 cover，返回 2 代表 root 上放置了摄像头。
func setCamera(root *TreeNode, hasParent bool, res *int) int {
	if root == nil {
		return -1
	}
	// 获取左右子节点的情况
	left := setCamera(root.Left, true, res)
	right := setCamera(root.Right, true, res)

	// 根据左右子节点的情况和父节点的情况判断当前节点应该做的事情
	if left == -1 && right == -1 {
		// 当前节点是叶子节点
		if hasParent {
			// 有父节点的话，让父节点来 cover 自己
			return 0
		}
		// 没有父节点的话，自己 set 一个摄像头
		*res++
		return 2
	}

	if left == 0 || right == 0 {
		// 左右子树存在没有被 cover 的
		// 必须在当前节点 set 一个摄像头
		*res++
		return 2
	}

	if left == 2 || right == 2 {
		// 左右子树只要有一个 set 了摄像头
		// 当前节点就已经是 cover 状态了
		return 1
	}

	// 剩下 left == 1 && right == 1 的情况
	// 即当前节点的左右子节点都被 cover
	if hasParent {
		// 如果有父节点的话，可以等父节点 cover 自己
		return 0
	} else {
		// 没有父节点，只能自己 set 一个摄像头
		*res++
		return 2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeCameras(t *testing.T) {

}
