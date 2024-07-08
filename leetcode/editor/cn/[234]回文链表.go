package leetcode

import (
	"testing"
)

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



 示例 1：


输入：head = [1,2,2,1]
输出：true


 示例 2：


输入：head = [1,2]
输出：false




 提示：


 链表中节点数目在范围[1, 10⁵] 内
 0 <= Node.val <= 9




 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

 Related Topics 栈 递归 链表 双指针 👍 1923 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//先用双指针找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	//反转后一半链表
	left, right := head, reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPalindromeLinkedList(t *testing.T) {

}
