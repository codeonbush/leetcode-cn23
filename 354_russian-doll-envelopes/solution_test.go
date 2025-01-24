package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

 注意：不允许旋转信封。

 示例 1：


输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

 示例 2：


输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1




 提示：


 1 <= envelopes.length <= 10⁵
 envelopes[i].length == 2
 1 <= wi, hi <= 10⁵


 Related Topics 数组 二分查找 动态规划 排序 👍 1077 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
import "sort"

// envelopes = [[w, h], [w, h]...]
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	// 按宽度升序排列，如果宽度一样，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 对高度数组寻找 LIS
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}

	return lengthOfLIS(height)
}

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))
	// 牌堆数初始化为 0
	var piles int
	for i := 0; i < len(nums); i++ {
		// 要处理的扑克牌
		poker := nums[i]

		// ***** 搜索左侧边界的二分查找 *****
		var left, right int = 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// ********************************

		// 没找到合适的牌堆，新建一堆
		if left == piles {
			piles++
		}
		// 把这张牌放到牌堆顶
		top[left] = poker
	}
	// 牌堆数就是 LIS 长度
	return piles
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRussianDollEnvelopes(t *testing.T) {
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
