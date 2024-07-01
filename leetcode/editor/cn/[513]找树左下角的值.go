package leetcode

import (
	"testing"
)

/**
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

 假设二叉树中至少有一个节点。



 示例 1:




输入: root = [2,1,3]
输出: 1


 示例 2:




输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7




 提示:


 二叉树的节点个数的范围是 [1,10⁴]

 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 577 👎 0

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
func findBottomLeftValue(root *TreeNode) int {
	maxDepth := 0 //最大深度
	res := 0
	pathD := 0 //当前深度
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		pathD++
		if pathD > maxDepth {
			maxDepth = pathD
			res = root.Val
		}

		traverse(root.Left)
		traverse(root.Right)
		pathD--
	}
	traverse(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindBottomLeftTreeValue(t *testing.T) {

}
