package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。



 示例 1：


输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]


 示例 2：


输入：nums = [4,4,3,2,1]
输出：[[4,4]]




 提示：


 1 <= nums.length <= 15
 -100 <= nums[i] <= 100


 Related Topics 位运算 数组 哈希表 回溯 👍 835 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack(nums, 0, &track, &res)
	return res
}

// 记录回溯的路径
// 回溯算法主函数
func backtrack(nums []int, start int, track *[]int, res *[][]int) {
	if len(*track) >= 2 {
		// 找到一个合法答案
		temp := make([]int, len(*track))
		copy(temp, *track)
		*res = append(*res, temp)
	}
	// 用哈希集合防止重复选择相同元素
	used := make(map[int]bool)
	// 回溯算法标准框架
	for i := start; i < len(nums); i++ {
		// 保证集合中元素都是递增顺序
		if len(*track) > 0 && (*track)[len(*track)-1] > nums[i] {
			continue
		}
		// 保证不要重复使用相同的元素
		if used[nums[i]] {
			continue
		}
		// 选择 nums[i]
		used[nums[i]] = true
		*track = append(*track, nums[i])
		// 递归遍历下一层回溯树
		backtrack(nums, i+1, track, res)
		// 撤销选择 nums[i]
		*track = (*track)[:len(*track)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNonDecreasingSubsequences(t *testing.T) {
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
