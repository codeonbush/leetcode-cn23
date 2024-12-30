package leetcode

import(
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/**
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。 

 candidates 中的每个数字在每个组合中只能使用 一次 。 

 注意：解集不能包含重复的组合。 

 

 示例 1: 

 
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
] 

 示例 2: 

 
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
] 

 

 提示: 

 
 1 <= candidates.length <= 100 
 1 <= candidates[i] <= 50 
 1 <= target <= 30 
 

 Related Topics 数组 回溯 👍 1614 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	// 记录 track 中的元素之和
	var trackSum int

	// 先排序，让相同的元素靠在一起
	sort.Ints(candidates)
	backtrack(candidates, target, 0, &track, &res, &trackSum)

	return res
}

// 回溯算法主函数
func backtrack(nums []int, target, start int, track *[]int, res *[][]int, trackSum *int) {
	// base case，达到目标和，找到符合条件的组合
	if *trackSum == target {
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return
	}
	// base case，超过目标和，直接结束
	if *trackSum > target {
		return
	}

	// 回溯算法标准框架
	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		// 做选择
		*track = append(*track, nums[i])
		*trackSum += nums[i]
		// 递归遍历下一层回溯树
		backtrack(nums, target, i+1, track, res, trackSum)
		// 撤销选择
		*track = (*track)[:len(*track)-1]
		*trackSum -= nums[i]
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func TestCombinationSumIi(t *testing.T){
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