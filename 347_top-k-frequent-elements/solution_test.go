package leetcode

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



 示例 1:


输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]


 示例 2:


输入: nums = [1], k = 1
输出: [1]



 提示：


 1 <= nums.length <= 10⁵
 k 的取值范围是 [1, 数组中不相同的元素的个数]
 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的




 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

 Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1889 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 用优先级队列解决这道题
func topKFrequent(nums []int, k int) []int {
	// nums 中的元素 -> 该元素出现的频率
	valToFreq := make(map[int]int)
	for _, v := range nums {
		valToFreq[v]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for key, value := range valToFreq {
		heap.Push(pq, &Element{key, value})
		if pq.Len() > k {
			// 弹出最小元素，维护队列内是 k 个频率最大的元素
			heap.Pop(pq)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		// res 数组中存储前 k 个最大元素
		res[i] = heap.Pop(pq).(*Element).value
	}

	return res
}

type Element struct {
	value    int
	priority int
}

type PriorityQueue []*Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 队列按照键值对中的值（元素出现频率）从小到大排序
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Element))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 用计数排序的方法解决这道题
func topKFrequent2(nums []int, k int) []int {
	// nums 中的元素 -> 该元素出现的频率
	valToFreq := make(map[int]int)
	for _, v := range nums {
		valToFreq[v]++
	}

	// 频率 -> 这个频率有哪些元素
	freqToVals := make([][]int, len(nums)+1)
	for val, freq := range valToFreq {
		freqToVals[freq] = append(freqToVals[freq], val)
	}

	res := make([]int, k)
	p := 0
	// freqToVals 从后往前存储着出现最多的元素
	for i := len(freqToVals) - 1; i > 0; i-- {
		for _, val := range freqToVals[i] {
			// 将出现次数最多的 k 个元素装入 res
			res[p] = val
			p++
			if p == k {
				return res
			}
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTopKFrequentElements(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{10, 6, 8, 5, 11, 9},
			expected: []int{3, 1, 1, 1, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := longestWPI(test.input)
			fmt.Println(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For heights %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
