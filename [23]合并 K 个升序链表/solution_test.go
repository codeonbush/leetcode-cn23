package leetcode

import (
	"container/heap"
	"testing"
)

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
// 示例 2：
//
// 输入：lists = []
// 输出：[]
//
// 示例 3：
//
// 输入：lists = [[]]
// 输出：[]
//
// 提示：
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2728 👎 0
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

// PriorityQueue 用于实现堆操作的 ListNode 列表。
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// mergeKLists 合并多个排序链表。
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 将所有链表的头节点加入最小堆。
	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}

	// 当优先队列不为空时，从堆中弹出最小元素，并将其添加到结果链表。
	for pq.Len() > 0 {
		smallest := heap.Pop(&pq).(*ListNode)
		current.Next = smallest
		current = current.Next
		if smallest.Next != nil {
			heap.Push(&pq, smallest.Next)
		}
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeKSortedLists(t *testing.T) {
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
