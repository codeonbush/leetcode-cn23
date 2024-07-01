package leetcode

import (
	"math"
	"testing"
)

/**
给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

 （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）



 示例 1：




输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
输出：7
解释：
我们有大量的节点与其祖先的差值，其中一些如下：
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。


 示例 2：


输入：root = [1,null,2,null,0,3]
输出：3




 提示：


 树中的节点数在 2 到 5000 之间。
 0 <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 二叉树 👍 266 👎 0

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
//每个节点需要知道左右子树的最小值和最大值，然后就能算出「以自己为祖先」的最大差值。
//
//每个节点都知道以自己为祖先的最大差值，那么所有这些差值中最大的那个就是整棵树的最大差值，
//这个取最大值的过程需要在后序遍历的位置进行
//注意到返回值使用[2]int要比[]int快
//在这个方法中，计算最大差值确实没有使用绝对值，但仍然能够正确地计算结果。这是因为通过分别计算 `rootMax - root.Val` 和 `root.Val - rootMin`，方法已经涵盖了两种情况：
//
//1. `rootMax - root.Val` 表示较大节点的值与当前根节点值之间的差，确保覆盖所有可能的较大节点在根节点之上的情况。
//2. `root.Val - rootMin` 表示当前根节点值与较小节点值之间的差，确保覆盖所有可能的较小节点在根节点之下的情况。
//
//将两者取最大值（`max1(res, rootMax - root.Val, root.Val - rootMin)`）已涵盖所有情况下的大差值，而无需显式计算绝对值。因此，求解最大祖先-节点间差值时，直接比较节点及其祖先间的最大和最小值差即可。
func maxAncestorDiff(root *TreeNode) int {
	res := 0

	// 定义：输入一棵二叉树，返回该二叉树中节点的最小值和最大值，
	// 第一个元素是最小值，第二个值是最大值
	var getMinMax func(*TreeNode) [2]int
	getMinMax = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{math.MaxInt32, math.MinInt32}
		}
		leftMinMax := getMinMax(root.Left)
		rightMinMax := getMinMax(root.Right)
		// 以 root 为根的这棵树的最大值和最小值可以通过左右子树的最大最小值推导出来
		rootMin := min(root.Val, leftMinMax[0], rightMinMax[0])
		rootMax := max1(root.Val, leftMinMax[1], rightMinMax[1])
		// 在后序位置顺便判断所有差值的最大值
		res = max1(res, rootMax-root.Val, root.Val-rootMin)

		return [2]int{rootMin, rootMax}
	}

	getMinMax(root)
	return res
}

func min(a, b, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

func max1(a, b, c int) int {
	return int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumDifferenceBetweenNodeAndAncestor(t *testing.T) {

}
