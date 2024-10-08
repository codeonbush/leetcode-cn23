package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，
并以数组的形式返回结果。



 示例 1：


输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]


 示例 2：


输入：nums = [1,1]
输出：[2]




 提示：


 n == nums.length
 1 <= n <= 10⁵
 1 <= nums[i] <= n


 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。

 Related Topics 数组 哈希表 👍 1333 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)
	for _, num := range nums {
		count[num]++
	}
	res := []int{}
	for num := 1; num <= n; num++ {
		if count[num] == 0 {
			res = append(res, num)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllNumbersDisappearedInAnArray(t *testing.T) {
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
