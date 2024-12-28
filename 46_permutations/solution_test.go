package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



 示例 1：


输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


 示例 2：


输入：nums = [0,1]
输出：[[0,1],[1,0]]


 示例 3：


输入：nums = [1]
输出：[[1]]




 提示：


 1 <= nums.length <= 6
 -10 <= nums[i] <= 10
 nums 中的所有整数 互不相同


 Related Topics 数组 回溯 👍 2954 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	return res
}

// 路径：记录在 track 中
// 选择列表：nums 中不存在于 track 的那些元素
// 结束条件：nums 中的元素全都在 track 中出现
func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	// 触发结束条件
	if len(track) == len(nums) {
		// 因为 track 是全局变量，因此需要新建一个数组来存储一份全排列
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := range nums {
		// 排除不合法的选择
		if used[i] {
			// 剪枝，避免重复使用同一个数字
			continue
		}
		// 做选择
		track = append(track, nums[i])
		used[i] = true
		// 进入下一层决策树
		backtrack(nums, track, used, res)
		// 取消选择
		track = track[:len(track)-1]
		used[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutations(t *testing.T) {
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
