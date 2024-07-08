package leetcode

import (
	"container/heap"
	"testing"
)

/**
给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。



 示例 1:


输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]


 示例 2:


输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]




 提示:


 1 <= nums1.length, nums2.length <= 10⁵
 -10⁹ <= nums1[i], nums2[i] <= 10⁹
 nums1 和 nums2 均为 升序排列

 1 <= k <= 10⁴
 k <= nums1.length * nums2.length


 Related Topics 数组 堆（优先队列） 👍 601 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 怎么把这道题变成合并多个有序链表呢？就比如说题目输入的用例：
//
// nums1 = [1,7,11], nums2 = [2,4,6]
// 组合出的所有数对儿这就可以抽象成三个有序链表：
//
// [1, 2] -> [1, 4] -> [1, 6]
// [7, 2] -> [7, 4] -> [7, 6]
// [11, 2] -> [11, 4] -> [11, 6]
// 这三个链表中每个元素（数对之和）是递增的，所以就可以按照 23. 合并K个升序链表 的思路来合并，取出前 k 个作为答案即可。
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 存储三元组 (num1[i], nums2[i], i)
	// i 记录 nums2 元素的索引位置，用于生成下一个节点
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 按照 23 题的逻辑初始化优先级队列
	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}

	res := [][]int{}
	// 执行合并多个有序链表的逻辑
	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		k--
		// 链表中的下一个节点加入优先级队列
		//数组中的第三个值为索引，索引+1
		nextIndex := cur[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})
		}

		res = append(res, []int{cur[0], cur[1]})
	}
	return res
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 按照数对的元素和升序排序
	return (pq[i][0] + pq[i][1]) < (pq[j][0] + pq[j][1])
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindKPairsWithSmallestSums(t *testing.T) {

}
