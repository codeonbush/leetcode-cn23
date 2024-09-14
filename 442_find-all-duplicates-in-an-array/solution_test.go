package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整
数，并以数组形式返回。

 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间（不包括存储输出所需的空间）的算法解决此问题。



 示例 1：


输入：nums = [4,3,2,7,8,2,3,1]
输出：[2,3]


 示例 2：


输入：nums = [1,1,2]
输出：[1]


 示例 3：


输入：nums = [1]
输出：[]




 提示：


 n == nums.length
 1 <= n <= 10⁵
 1 <= nums[i] <= n
 nums 中的每个元素出现 一次 或 两次


 Related Topics 数组 哈希表 👍 793 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findDuplicates(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	// 用数组模拟哈希集合
	seen := make([]int, n+1)
	for _, num := range nums {
		if seen[num] == 0 {
			// 添加到哈希集合
			seen[num] = 1
		} else {
			// 找到重复元素
			res = append(res, num)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllDuplicatesInAnArray(t *testing.T) {
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
