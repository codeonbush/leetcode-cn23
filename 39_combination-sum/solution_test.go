package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所
有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。 

 candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 

 对于给定的输入，保证和为 target 的不同组合数少于 150 个。 

 

 示例 1： 

 
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。 

 示例 2： 

 
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]] 

 示例 3： 

 
输入: candidates = [2], target = 1
输出: []
 

 

 提示： 

 
 1 <= candidates.length <= 30 
 2 <= candidates[i] <= 40 
 candidates 的所有元素 互不相同 
 1 <= target <= 40 
 

 Related Topics 数组 回溯 👍 2944 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	// 记录回溯的路径
	var track []int
	// 记录 track 中的路径和
	trackSum := 0

	var backtrack func([]int, int)
	backtrack = func(nums []int, start int) {
		// base case，找到目标和，记录结果
		if trackSum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		// base case，超过目标和，停止向下遍历
		if trackSum > target {
			return
		}
		// 回溯算法标准框架
		for i := start; i < len(nums); i++ {
			// 选择 nums[i]
			trackSum += nums[i]
			track = append(track, nums[i])
			// 递归遍历下一层回溯树
			// 同一元素可重复使用，注意参数
			backtrack(nums, i)
			// 撤销选择 nums[i]
			trackSum -= nums[i]
			track = track[:len(track)-1]
		}
	}

	backtrack(candidates, 0)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func TestCombinationSum(t *testing.T){
    tests := []struct {
		input  []int
		expected []int
	}{
		{
			input:  []int{10, 6, 8, 5, 11, 9},
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