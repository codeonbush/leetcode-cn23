package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。 

 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列
。 

 示例 1： 

 
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 

 示例 2： 

 
输入：nums = [0,1,0,3,2,3]
输出：4
 

 示例 3： 

 
输入：nums = [7,7,7,7,7,7,7]
输出：1
 

 

 提示： 

 
 1 <= nums.length <= 2500 
 -10⁴ <= nums[i] <= 10⁴ 
 

 

 进阶： 

 
 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 
 

 Related Topics 数组 二分查找 动态规划 👍 3844 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	// 定义：dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	// base case：dp 数组全都初始化为 1
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// dp[i] = Math.max(dp[i], dp[j] + 1);
				dp[i] = max(dp[i], dp[j] + 1);
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)


func TestLongestIncreasingSubsequence(t *testing.T){
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