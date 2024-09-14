package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

 请实现 KthLargest 类：


 KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
 int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。




 示例 1：


 输入： ["KthLargest", "add", "add", "add", "add", "add"] [[3, [4, 5, 8, 2]], [3],
[5], [10], [9], [4]]


 输出：[null, 4, 5, 5, 8, 8]

 解释：

 KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]); kthLargest.add(3); //
返回 4 kthLargest.add(5); // 返回 5 kthLargest.add(10); // 返回 5 kthLargest.add(9); /
/ 返回 8 kthLargest.add(4); // 返回 8



 示例 2：


 输入： ["KthLargest", "add", "add", "add", "add"] [[4, [7, 7, 7, 7, 8, 3]], [2], [
10], [9], [9]]


 输出：[null, 7, 7, 7, 8]

 解释： KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
 kthLargest.add(2); // 返回 7
 kthLargest.add(10); // 返回 7
 kthLargest.add(9); // 返回 7
 kthLargest.add(9); // 返回 8


提示：


 0 <= nums.length <= 10⁴
 1 <= k <= nums.length + 1
 -10⁴ <= nums[i] <= 10⁴
 -10⁴ <= val <= 10⁴
 最多调用 add 方法 10⁴ 次


 Related Topics 树 设计 二叉搜索树 二叉树 数据流 堆（优先队列） 👍 480 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	k  int
	pq *PriorityQueue
}

// 默认是小顶堆
type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	pq := &PriorityQueue{}
	heap.Init(pq)
	// 将 nums 装入小顶堆，保留下前 k 大的元素
	for _, e := range nums {
		heap.Push(pq, e)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	return KthLargest{k: k, pq: pq}
}

func (this *KthLargest) Add(val int) int {
	// 维护小顶堆只保留前 k 大的元素
	heap.Push(this.pq, val)
	if this.pq.Len() > this.k {
		heap.Pop(this.pq)
	}
	// 堆顶就是第 k 大元素（即倒数第 k 小的元素）
	return (*this.pq)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAStream(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{10, 6, 8, 5, 11, 9},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			kt := Constructor(3, test.input)
			fmt.Println(kt.Add(1))
			fmt.Println(kt.Add(2))
			fmt.Println(kt.Add(3))
			fmt.Println(kt.Add(4))
		})
	}
}
