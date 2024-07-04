package leetcode

import (
	"testing"
)

/**
给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。



 示例 1:




输入: head = [-10,-3,0,5,9]
输出: [0,-3,9,-10,null,5]
解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。


 示例 2:


输入: head = []
输出: []




 提示:


 head 中的节点数在[0, 2 * 10⁴] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 904 👎 0

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
// 通过中序遍历特点写出的解法
func sortedListToBST(head *ListNode) *TreeNode {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	return inorderBuild(&head, 0, l-1)
}

func inorderBuild(head **ListNode, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	// 构造左子树
	leftTree := inorderBuild(head, left, mid-1)
	// 构造根节点
	root := &TreeNode{Val: (*head).Val}
	*head = (*head).Next
	// 构造右子树
	rightTree := inorderBuild(head, mid+1, right)
	// 将左右子树接到根节点上
	root.Left = leftTree
	root.Right = rightTree
	return root
}

//// 解法二、通过找链表中点的方式写出的解法
//func sortedListToBST(head *ListNode) *TreeNode {
//   return build(head, nil)
//}
//
//// 把链表左闭右开区间 [begin, end) 的节点构造成 BST
//func build(begin, end *ListNode) *TreeNode {
//	if begin == end {
//		// 因为是左闭右开区间，所以现在已经成空集了
//		return nil
//	}
//	mid := getMid(begin, end)
//	// 构造根节点
//	root := &TreeNode{Val: mid.Val}
//	// 构造左子树
//	root.Left = build(begin, mid)
//	// 构造右子树
//	root.Right = build(mid.Next, end)
//	return root
//}
//
//// 获取链表左闭右开区间 [begin, end) 的中心节点
//func getMid(begin, end *ListNode) *ListNode {
//	slow, fast := begin, begin
//	for fast != end && fast.Next != end {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return slow
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertSortedListToBinarySearchTree(t *testing.T) {

}
