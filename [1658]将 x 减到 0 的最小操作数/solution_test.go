package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数
组以供接下来的操作使用。

 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。



 示例 1：


输入：nums = [1,1,4,2,3], x = 5
输出：2
解释：最佳解决方案是移除后两个元素，将 x 减到 0 。


 示例 2：


输入：nums = [5,6,7,8,9], x = 4
输出：-1


 示例 3：


输入：nums = [3,2,20,1,1,3], x = 10
输出：5
解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。




 提示：


 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁴
 1 <= x <= 10⁹


 Related Topics 数组 哈希表 二分查找 前缀和 滑动窗口 👍 358 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 滑动窗口需要寻找的子数组目标和
	target := sum - x

	left, right := 0, 0
	// 记录窗口内所有元素和
	windowSum := 0
	// 记录目标子数组的最大长度
	maxLen := -1
	// 开始执行滑动窗口框架
	for right < len(nums) {
		// 扩大窗口
		windowSum += nums[right]
		right++

		for windowSum > target && left < right {
			// 缩小窗口
			windowSum -= nums[left]
			left++
		}

		// 寻找目标子数组, 更新最大长度
		if windowSum == target {
			if maxLen == -1 || right-left > maxLen {
				maxLen = right - left
			}
		}
	}
	// 目标子数组的最大长度可以推导出需要删除的字符数量
	if maxLen == -1 {
		return -1
	}
	return n - maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumOperationsToReduceXToZero(t *testing.T) {

}
