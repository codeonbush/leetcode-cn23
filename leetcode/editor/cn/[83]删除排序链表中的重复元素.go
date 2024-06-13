package leetcode

import (
	"testing"
)

// 给定一个已排序的链表的头
// head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
// 示例 1：
//
// 输入：head = [1,1,2]
// 输出：[1,2]
//
// 示例 2：
//
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]
//
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
//
// Related Topics 链表 👍 1117 👎 0

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			// slow指针与快指针相等的值赋值为快指针的值
			slow.Next = fast
			// 将指针后移，操作下一个节点
			slow = slow.Next
		}
		// 将指针后移，操作下一个节点
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {

}
