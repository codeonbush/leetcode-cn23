package leetcode

import (
	"testing"
)

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
//
//
// Related Topics 链表 双指针 👍 1282 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Iterative solution
// 和83区别在于，把重复元素全部删除一个不留
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p, q := dummy, head
	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			// 发现重复节点，跳过这些重复节点
			for q.Next != nil && q.Val == q.Next.Val {
				q = q.Next
			}
			q = q.Next
			// 此时 q 跳过了这一段重复元素
			if q == nil {
				p.Next = nil
			}
			// 不过下一段元素也可能重复，等下一轮 while 循环判断
		} else {
			// 不是重复节点，接到 dummy 后面
			p.Next = q
			p = p.Next
			q = q.Next
		}
	}
	return dummy.Next
}

// 递归解法
func deleteDuplicatesRecursive(head *ListNode) *ListNode {
	// 定义：输入一条单链表头结点，返回去重之后的单链表头结点
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		// 如果头结点和身后节点的值不同，则对之后的链表去重即可
		head.Next = deleteDuplicatesRecursive(head.Next)
		return head
	}
	// 如果如果头结点和身后节点的值相同，则说明从 head 开始存在若干重复节点
	// 越过重复节点，找到 head 之后那个不重复的节点
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	// 直接返回那个不重复节点开头的链表的去重结果，就把重复节点删掉了
	return deleteDuplicatesRecursive(head.Next)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedListIi(t *testing.T) {

}
