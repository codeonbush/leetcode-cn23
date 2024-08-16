package leetcode

import (
	"container/heap"
	"testing"
)

/**
给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。 请注意，它是 排序后 的第 k 小元素，而不是第 k
个 不同 的元素。

 你必须找到一个内存复杂度优于 O(n²) 的解决方案。



 示例 1：


输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
输出：13
解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13


 示例 2：


输入：matrix = [[-5]], k = 1
输出：-5




 提示：


 n == matrix.length
 n == matrix[i].length
 1 <= n <= 300
 -10⁹ <= matrix[i][j] <= 10⁹
 题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
 1 <= k <= n²




 进阶：


 你能否用一个恒定的内存(即 O(1) 内存复杂度)来解决这个问题?
 你能在 O(n) 的时间复杂度下解决这个问题吗?这个方法对于面试来说可能太超前了，但是你会发现阅读这篇文章（ this paper ）很有趣。


 Related Topics 数组 二分查找 矩阵 排序 堆（优先队列） 👍 1059 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 矩阵中的每一行都是排好序的，就好比多条有序链表，
// 你用优先级队列施展合并多条有序链表的逻辑就能找到第 k 小的元素了。
func kthSmallest(matrix [][]int, k int) int {
	// 自定义一个最小堆类型
	pq := IntHeap{}
	// 初始化堆，把每一行的第一个元素装进去
	for i := 0; i < len(matrix); i++ {
		pq = append(pq, Item{value: matrix[i][0], row: i, col: 0})
	}
	heap.Init(&pq)

	var res int
	// 执行合并多个有序链表的逻辑，找到第 k 小的元素
	for k > 0 && pq.Len() > 0 {
		cur := heap.Pop(&pq).(Item)
		res = cur.value
		k--
		// 链表中的下一个节点加入堆
		row, col := cur.row, cur.col+1
		if col < len(matrix[row]) {
			heap.Push(&pq, Item{value: matrix[row][col], row: row, col: col})
		}
	}

	return res
}

// 定义一个 Item 类型，表示堆中的元素
type Item struct {
	value int // 当前元素的值
	row   int // 当前元素所在的行
	col   int // 当前元素所在的列
}

// 定义一个最小堆类型 IntHeap
// 实现 heap.Interface 接口的方法
type IntHeap []Item

func (t IntHeap) Len() int {
	return len(t)
}

func (t IntHeap) Less(i, j int) bool {
	return t[i].value < t[j].value
}

func (t IntHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *IntHeap) Push(x interface{}) {
	*t = append(*t, x.(Item))
}

func (t *IntHeap) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthSmallestElementInASortedMatrix(t *testing.T) {

}
