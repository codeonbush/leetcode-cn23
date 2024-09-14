package leetcode

import (
	"container/heap"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/**
给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。

 返回 已排序的字符串 。如果有多个答案，返回其中任何一个。



 示例 1:


输入: s = "tree"
输出: "eert"
解释: 'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。


 示例 2:


输入: s = "cccaaa"
输出: "cccaaa"
解释: 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
注意"cacaca"是不正确的，因为相同的字母必须放在一起。


 示例 3:


输入: s = "Aabb"
输出: "bbAa"
解释: 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
注意'A'和'a'被认为是两种不同的字符。




 提示:


 1 <= s.length <= 5 * 10⁵
 s 由大小写英文字母和数字组成


 Related Topics 哈希表 字符串 桶排序 计数 排序 堆（优先队列） 👍 527 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	chars := []rune(s)
	// s 中的字符 -> 该字符出现的频率
	charToFreq := make(map[rune]int)
	for _, ch := range chars {
		charToFreq[ch]++
	}

	pq := make(PriorityQueue, 0, len(charToFreq))
	// 按照字符频率排序
	for entry := range charToFreq {
		heap.Push(&pq, &Item{
			value:    entry,
			priority: charToFreq[entry],
		})
	}
	heap.Init(&pq)

	var sb strings.Builder
	for pq.Len() > 0 {
		// 把频率最高的字符排在前面
		item := heap.Pop(&pq).(*Item)
		part := strings.Repeat(string(item.value), item.priority)
		sb.WriteString(part)
	}

	return sb.String()
}

type Item struct {
	value    rune
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 队列按照键值对中的值（字符出现频率）从大到小排序
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSortCharactersByFrequency(t *testing.T) {
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
