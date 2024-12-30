package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。 

 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 

 

 示例 1： 

 
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 

 示例 2： 

 
输入：nums = [0]
输出：[[],[0]]
 

 

 提示： 

 
 1 <= nums.length <= 10 
 -10 <= nums[i] <= 10 
 nums 中的所有元素 互不相同 
 

 Related Topics 位运算 数组 回溯 👍 2404 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var res [][]int
	// 记录回溯算法的递归路径
	var track []int

	// 回溯算法核心函数，遍历子集问题的回溯树
	var backtrack func(int)
	backtrack = func(start int) {
		// 前序位置，每个节点的值都是一个子集
		res = append(res, append([]int(nil), track...))

		// 回溯算法标准框架
		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集
			backtrack(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSubsets(t *testing.T){
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