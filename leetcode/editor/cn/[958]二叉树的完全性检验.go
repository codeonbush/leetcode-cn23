package leetcode

import (
	"testing"
)

/**
给你一棵二叉树的根节点
 root ，请你判断这棵树是否是一棵 完全二叉树 。

 在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含
 1 到
 2ʰ 个节点。



 示例 1：




输入：root = [1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6}）尽可能靠左。


 示例 2：




输入：root = [1,2,3,4,5,null,7]
输出：false
解释：值为 7 的节点不满足条件「节点尽可能靠左」。




 提示：


 树中节点数目在范围 [1, 100] 内
 1 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 289 👎 0

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
//完全二叉树最终队列留下的都是空指针
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	// 遍历完所有非空节点时变成 true
	end := false
	// while 循环控制从上向下一层层遍历
	for len(q) > 0 {
		sz := len(q)
		// for 循环控制每一层从左向右遍历
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == nil {
				// 第一次遇到 null 时 end 变成 true
				// 如果之后的所有节点都是 null，则说明是完全二叉树
				end = true
			} else {
				if end {
					// end 为 true 时遇到非空节点说明不是完全二叉树
					return false
				}
				// 将下一层节点放入队列，不用判断是否非空
				q = append(q, cur.Left, cur.Right)
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCheckCompletenessOfABinaryTree(t *testing.T) {

}
