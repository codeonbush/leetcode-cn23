package leetcode

import (
	"testing"
)

// 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//
// 如果有两个中间结点，则返回第二个中间结点。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5]
// 输出：[3,4,5]
// 解释：链表只有一个中间结点，值为 3 。
//
// 示例 2：
//
// 输入：head = [1,2,3,4,5,6]
// 输出：[4,5,6]
// 解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。
//
// 提示：
//
// 链表的结点数范围是 [1, 100]
// 1 <= Node.val <= 100
//
// Related Topics 链表 双指针 👍 969 👎 0
type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 让两个指针 slow 和 fast 分别指向链表头结点 head。
// 每当慢指针 slow 前进一步，快指针 fast 就前进两步，这样，当 fast 走到链表末尾时，slow 就指向了链表中点。
func middleNode(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return p2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMiddleOfTheLinkedList(t *testing.T) {
	//cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d expected %d, but %d got",
	//				c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}
}
