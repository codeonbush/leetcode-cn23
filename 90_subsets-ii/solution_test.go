package leetcode

import(
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。 

 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 

 
 
 
 
 

 示例 1： 

 
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
 

 示例 2： 

 
输入：nums = [0]
输出：[[],[0]]
 

 

 提示： 

 
 1 <= nums.length <= 10 
 -10 <= nums[i] <= 10 
 

 Related Topics 位运算 数组 回溯 👍 1255 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	backtrack(nums, &track, &res, 0)

	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, start int) {
	// 前序位置，每个节点的值都是一个子集
	tmp := make([]int, len(*track))
	copy(tmp, *track)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*track = append(*track, nums[i])
		backtrack(nums, track, res, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSubsetsIi(t *testing.T){
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