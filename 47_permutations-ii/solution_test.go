package leetcode

import(
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/**
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。 

 

 示例 1： 

 
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
 

 示例 2： 

 
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 

 

 提示： 

 
 1 <= nums.length <= 8 
 -10 <= nums[i] <= 10 
 

 Related Topics 数组 回溯 👍 1640 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))

	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	backtrack(nums, &res, track, used)
	return res
}

func backtrack(nums []int, res *[][]int, track []int, used []bool) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 新添加的剪枝逻辑，固定相同的元素在排列中的相对位置
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, res, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func TestPermutationsIi(t *testing.T){
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