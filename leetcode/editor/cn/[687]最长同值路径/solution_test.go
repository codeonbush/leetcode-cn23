package leetcode

import (
	"testing"
)

/**
给定一个二叉树的
 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。

 两个节点之间的路径长度 由它们之间的边数表示。



 示例 1:




输入：root = [5,4,5,1,1,5]
输出：2


 示例 2:




输入：root = [1,4,5,4,4,5]
输出：2




 提示:


 树的节点数的范围是
 [0, 10⁴]
 -1000 <= Node.val <= 1000
 树的深度将不超过 1000


 Related Topics 树 深度优先搜索 二叉树 👍 805 👎 0

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
//不要把路径长度和节点个数混淆，路径长度=节点个数-1
//最长同值路径
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	// 在后序遍历的位置更新 res
	maxLen(root, root.Val, &res)
	return res
}

// 计算以 root 为根的这棵二叉树中，从 root 开始值为 parentVal 的最长树枝长度
func maxLen(root *TreeNode, parentVal int, res *int) int {
	if root == nil {
		return 0
	}
	//利用函数定义，计算左右子树值为 root.val 的最长树枝长度
	leftLen := maxLen(root.Left, root.Val, res)
	rightLen := maxLen(root.Right, root.Val, res)
	//后序遍历位置顺便更新全局变量
	//同值路径就是左右同值树枝长度之和
	//不要把路径长度和节点个数混淆，路径长度=节点个数-1，左右长度已经隐含了当前节点，所以不需要+1
	if *res < leftLen+rightLen {
		*res = leftLen + rightLen
	}
	//如果 root 本身和上级值不同，那么整棵子树都不可能有同值树枝
	if root.Val != parentVal {
		return 0
	}
	//实现函数的定义：
	//以 root 为根的二叉树从 root 开始值为 parentVal 的最长树枝长度
	//等于左右子树的最长树枝长度的最大值加上 root 节点本身
	if leftLen > rightLen {
		return 1 + leftLen
	}
	return 1 + rightLen
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestUnivaluePath(t *testing.T) {

}
