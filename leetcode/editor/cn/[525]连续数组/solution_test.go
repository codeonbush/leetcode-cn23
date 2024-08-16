package leetcode

import (
	"testing"
)

/**
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。



 示例 1:


输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

 示例 2:


输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。



 提示：


 1 <= nums.length <= 10⁵
 nums[i] 不是 0 就是 1


 Related Topics 数组 哈希表 前缀和 👍 742 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	// 通过一个值映射索引的哈希表以减少空间复杂度
	valToIndex := make(map[int]int)
	valToIndex[0] = -1 // 基本条件初始化

	// 保存前缀和
	preSum, res := 0, 0
	for i, num := range nums {
		// 使用一次累加(prefix sum思想)
		if num == 0 {
			preSum--
		} else {
			preSum++
		}

		// 若该前缀和已经在哈希表中，则更新res，使其为最大长度
		if idx, exists := valToIndex[preSum]; exists {
			res = max(res, i-idx)
		} else {
			valToIndex[preSum] = i // 新的前缀和，记录其索引
		}
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

func TestContiguousArray(t *testing.T) {

}
