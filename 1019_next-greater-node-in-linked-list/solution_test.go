package leetcode

import (
	. "cn"
	"testing"
)

/**
给定一个长度为 n 的链表 head

 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。

 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设
置 answer[i] = 0 。



 示例 1：




输入：head = [2,1,5]
输出：[5,5,0]


 示例 2：




输入：head = [2,7,4,3,5]
输出：[7,0,5,5,0]




 提示：


 链表中节点数为 n
 1 <= n <= 10⁴
 1 <= Node.val <= 10⁹


 Related Topics 栈 数组 链表 单调栈 👍 334 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	// 把单链表转化成数组，方便通过索引访问
	nums := []int{}
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	// 存放答案的数组
	res := make([]int, len(nums))
	stk := []int{}
	// 单调栈模板，求下一个更大元素，从后往前遍历
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= nums[i] {
			stk = stk[:len(stk)-1]
		}
		// 本题要求没有下一个更大元素时返回 0
		res[i] = 0
		if len(stk) > 0 {
			res[i] = stk[len(stk)-1]
		}
		stk = append(stk, nums[i])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterNodeInLinkedList(t *testing.T) {

}
