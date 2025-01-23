package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断
 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。

 如果你能使这个正方形，则返回 true ，否则返回 false 。



 示例 1:




输入: matchsticks = [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。


 示例 2:


输入: matchsticks = [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。




 提示:


 1 <= matchsticks.length <= 15
 1 <= matchsticks[i] <= 10⁸


 Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 550 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

func makesquare(matchsticks []int) bool {
	return canPartitionKSubsets(matchsticks, 4)
}

// 下面的代码需要学习前文 经典回溯算法：集合划分问题
// https://labuladong.online/algo/practice-in-action/partition-to-k-equal-sum-subsets/
func canPartitionKSubsets(nums []int, k int) bool {
	// 排除一些基本情况
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	// 使用位图技巧
	used := 0
	target := sum / k
	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	memo := make(map[int]bool)
	return backtrack(k, 0, nums, 0, used, target, memo)
}

func backtrack(k, bucket int, nums []int, start, used, target int, memo map[int]bool) bool {
	// base case
	if k == 0 {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		return true
	}
	if bucket == target {
		// 装满了当前桶，递归穷举下一个桶的选择
		// 让下一个桶从 nums[0] 开始选数字
		if res, ok := memo[used]; ok {
			// 缓存结果
			return res
		}
		res := backtrack(k-1, 0, nums, 0, used, target, memo)
		memo[used] = res
		return res
	}

	if res, ok := memo[used]; ok {
		// 避免冗余计算
		return res
	}

	for i := start; i < len(nums); i++ {
		// 剪枝
		// 判断第 i 位是否是 1
		if (used >> i & 1) == 1 {
			// nums[i] 已经被装入别的桶中
			continue
		}
		if nums[i]+bucket > target {
			continue
		}
		// 做选择
		// 将第 i 位置为 1
		used |= 1 << i
		bucket += nums[i]
		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucket, nums, i+1, used, target, memo) {
			return true
		}
		// 撤销选择
		// 使用异或运算将第 i 位恢复 0
		used ^= 1 << i
		bucket -= nums[i]
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMatchsticksToSquare(t *testing.T) {
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
