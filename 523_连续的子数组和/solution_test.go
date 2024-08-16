package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，如果 nums 有一个 好的子数组 返回 true ，否则返回 false：

 一个 好的子数组 是：


 长度 至少为 2 ，且
 子数组元素总和为 k 的倍数。


 注意：


 子数组 是数组中 连续 的部分。
 如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终 视为 k 的一个倍数。




 示例 1：


输入：nums = [23,2,4,6,7], k = 6
输出：true
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。

 示例 2：


输入：nums = [23,2,6,4,7], k = 6
输出：true
解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。


 示例 3：


输入：nums = [23,2,6,4,7], k = 13
输出：false




 提示：


 1 <= nums.length <= 10⁵
 0 <= nums[i] <= 10⁹
 0 <= sum(nums[i]) <= 2³¹ - 1
 1 <= k <= 2³¹ - 1


 Related Topics 数组 哈希表 数学 前缀和 👍 579 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	// 计算 nums 的前缀和
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	// 前缀和与 k 的余数到索引的映射，方便快速查找所需的前缀和
	valToIndex := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		// 在哈希表中记录余数
		val := preSum[i] % k
		// 如果这个余数还没有对应的索引，则记录下来
		if _, exists := valToIndex[val]; !exists {
			valToIndex[val] = i
		}
		// 如果这个前缀和已经有对应的索引了，则什么都不做
		// 因为题目想找长度最大的子数组，所以前缀和索引应尽可能小
	}
	for i := 1; i < len(preSum); i++ {
		// 计算 need，使得 (preSum[i] - need) % k == 0
		need := preSum[i] % k
		if idx, exists := valToIndex[need]; exists {
			if i-idx >= 2 {
				// 这个子数组的长度至少为 2
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContinuousSubarraySum(t *testing.T) {

}
