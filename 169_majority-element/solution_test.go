package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

 你可以假设数组是非空的，并且给定的数组总是存在多数元素。



 示例 1：


输入：nums = [3,2,3]
输出：3

 示例 2：


输入：nums = [2,2,1,1,1,2,2]
输出：2



提示：


 n == nums.length
 1 <= n <= 5 * 10⁴
 -10⁹ <= nums[i] <= 10⁹




 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

 Related Topics 数组 哈希表 分治 计数 排序 👍 2272 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	limit := len(nums) / 2
	for _, num := range nums {
		countMap[num]++
		if countMap[num] > limit {
			return num
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMajorityElement(t *testing.T) {
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
