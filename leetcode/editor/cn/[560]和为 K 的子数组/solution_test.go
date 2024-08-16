package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

 子数组是数组中元素的连续非空序列。



 示例 1：


输入：nums = [1,1,1], k = 2
输出：2


 示例 2：


输入：nums = [1,2,3], k = 3
输出：2




 提示：


 1 <= nums.length <= 2 * 10⁴
 -1000 <= nums[i] <= 1000
 -10⁷ <= k <= 10⁷


 Related Topics 数组 哈希表 前缀和 👍 2368 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	res := 0
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		preSum[i] = nums[i-1] + preSum[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if preSum[j+1]-preSum[i] == k {
				res++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubarraySumEqualsK(t *testing.T) {

}
