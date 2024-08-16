package leetcode

import (
	"testing"
)

/**
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。



 示例 1：


输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。


 示例 2：


输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。



 提示：


 树上节点的数目在范围 [2, 1000] 内
 -2³¹ <= Node.val <= 2³¹ - 1




 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 957 👎 0

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
//这道题需要用到「遍历」的思维模式。
//
//得益于二叉搜索树左小右大的特性，一个重要性质是：二叉搜索树的中序遍历结果有序。
//
//题目说有两个节点的值反了，也就是说中序遍历结果不再是有序的，有两个元素的位置反了。
//
//那么我们在中序遍历的时候找到破坏有序性的这两个元素，交换它们即可。
func recoverTree(root *TreeNode) {
	var first, second *TreeNode
	prev := &TreeNode{Val: -1 << 63}

	var inorderTraverse func(*TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraverse(root.Left)
		// 中序遍历代码位置，找出那两个节点
		if root.Val < prev.Val {
			// root 不符合有序性
			if first == nil {
				// 第一个错位节点是 prev
				first = prev
			}
			// 第二个错位节点是 root
			second = root
		}
		prev = root
		inorderTraverse(root.Right)
	}

	inorderTraverse(root)

	temp := first.Val
	first.Val = second.Val
	second.Val = temp
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRecoverBinarySearchTree(t *testing.T) {

}
