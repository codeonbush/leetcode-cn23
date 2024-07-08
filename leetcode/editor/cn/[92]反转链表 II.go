package leetcode

import (
	"testing"
)

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节
点，返回 反转后的链表 。



 示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]


 示例 2：


输入：head = [5], left = 1, right = 1
输出：[5]




 提示：


 链表中节点数目为 n
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n




 进阶： 你可以使用一趟扫描完成反转吗？

 Related Topics 链表 👍 1811 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// base case
	if m == 1 {
		return reverseN(head, n)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var successor *ListNode // 后继节点

// 反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// 以 head.Next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)

	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last

}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {

}
