package leetcode

import (
	"testing"
)

/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：


L0 → L1 → … → Ln - 1 → Ln


 请将其重新排列后变为：


L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



 示例 1：




输入：head = [1,2,3,4]
输出：[1,4,2,3]

 示例 2：




输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]



 提示：


 链表的长度范围为 [1, 5 * 10⁴]
 1 <= node.val <= 1000


 Related Topics 栈 递归 链表 双指针 👍 1496 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// 先把所有节点装进栈里，得到倒序结果
	type stack []*ListNode
	var stk stack
	p := head
	for p != nil {
		stk = append(stk, p)
		p = p.Next
	}

	p = head
	for p != nil {
		// 链表尾部的节点
		lastNode := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		next := p.Next
		if lastNode == next || lastNode.Next == next {
			// 结束条件，链表节点数为奇数或偶数时均适用
			lastNode.Next = nil
			break
		}
		p.Next = lastNode
		lastNode.Next = next
		p = next
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReorderList(t *testing.T) {

}
