package leetcode

import (
	"testing"
)

/**
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

 k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。



 示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]


 示例 2：




输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]



提示：


 链表中的节点数目为 n
 1 <= k <= n <= 5000
 0 <= Node.val <= 1000




 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？




 Related Topics 递归 链表 👍 2358 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前 k 个元素
	newHead := reverse(a, b)
	// 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k)

	return newHead
}

// 反转区间 [a, b) 的元素，注意是左闭右开
func reverse(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, a, a
	// while 终止的条件改一下就行了
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseNodesInKGroup(t *testing.T) {

}
