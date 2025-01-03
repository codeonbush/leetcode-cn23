package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。 

 

 示例 1： 

 
输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。 

 示例 2: 

 
输入: nums = [1,2,3,4], k = 3
输出: false 

 

 提示： 

 
 1 <= k <= len(nums) <= 16 
 0 < nums[i] < 10000 
 每个元素的频率在 [1,4] 范围内 
 

 Related Topics 位运算 记忆化搜索 数组 动态规划 回溯 状态压缩 👍 1095 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func canPartitionKSubsets(nums []int, k int) bool {
	// 排除一些基本情况
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum % k != 0 {
		// 不能整除，提前退出
		return false
	}
	// 使用位图技巧
	used := 0
	target := sum / k
	// 备忘录
	memo := make(map[int]bool)
	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	return backtrack(k, 0, nums, 0, used, target, memo)
}

func backtrack(k, bucket int, nums []int, start int, used int, target int, memo map[int]bool) bool {
	// base case
	if k == 0 {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		return true
	}
	if bucket == target {
		// 装满了当前桶，递归穷举下一个桶的选择
		// 让下一个桶从 nums[0] 开始选数字
		res, ok := memo[used]
		if !ok {
			res = backtrack(k-1, 0, nums, 0, used, target, memo)
			// 缓存结果
			memo[used] = res
		}
		return res
	}
	if res, ok := memo[used]; ok {
		// 避免冗余计算
		return res
	}
	for i := start; i < len(nums); i++ {
		// 剪枝
		if (used >> i) & 1 == 1 {
			// nums[i] 已经被装入别的桶中
			continue
		}
		if nums[i] + bucket > target {
			// 装不下，剪枝
			continue
		}
		// 做选择
		used |= 1 << i
		bucket += nums[i]
		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucket, nums, i+1, used, target, memo) {
			return true
		}
		// 撤销选择
		used ^= 1 << i
		bucket -= nums[i]
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)


func TestPartitionToKEqualSumSubsets(t *testing.T){
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