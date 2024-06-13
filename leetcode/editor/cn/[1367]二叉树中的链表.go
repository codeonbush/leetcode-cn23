package leetcode

import (
	"testing"
)

/**
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。


 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。



 示例 1：



 输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1
,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。


 示例 2：



 输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
null,1,3]
输出：true


 示例 3：

 输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。




 提示：


 二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
 链表包含的节点数目在 1 到 100 之间。
 二叉树包含的节点数目在 1 到 2500 之间。


 Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 189 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func isSubPath(head *ListNode, root *TreeNode) bool {
//	res := false
//	var traverse func(root *TreeNode, node *ListNode)
//	traverse = func(root *TreeNode, node *ListNode) {
//		if root == nil && node == nil {
//			res = true
//			return
//		}
//		if isSamePath(head, node) {
//			res = true
//			return
//		}
//		if root.Left != nil {
//			node.Next = &ListNode{Val: root.Left.Val}
//		}
//		traverse(root.Left, node)
//		traverse(root.Right, node)
//	}
//	traverse(root, &ListNode{})
//	return res
//}
//
//// 判断两条链表是否相同
//func isSamePath(p *ListNode, q *ListNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//	if p == nil || q == nil {
//		return false
//	}
//	if p.Val != q.Val {
//		return false
//	}
//	return isSamePath(p.Next, q.Next)
//}

// 本质上，isSubPath 就是在遍历二叉树的所有节点，对每个节点用 check 函数判断是否能够将链表嵌进去。
func isSubPath(head *ListNode, root *TreeNode) bool {
	// base case
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	// 当找到一个二叉树节点的值等于链表头结点时
	if head.Val == root.Val {
		// 判断是否能把链表嵌进去
		if check(head, root) {
			return true
		}
	}
	// 继续去遍历其他节点尝试嵌入链表
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// 检查是否能够将链表嵌入二叉树
func check(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	if head.Val == root.Val {
		// 在子树上嵌入子链表
		return check(head.Next, root.Left) || check(head.Next, root.Right)
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLinkedListInBinaryTree(t *testing.T) {

}
