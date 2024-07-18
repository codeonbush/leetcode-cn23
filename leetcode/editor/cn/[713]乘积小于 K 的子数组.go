package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。



 示例 1：


输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。


 示例 2：


输入：nums = [1,2,3], k = 0
输出：0



 提示:


 1 <= nums.length <= 3 * 10⁴
 1 <= nums[i] <= 1000
 0 <= k <= 10⁶


 Related Topics 数组 滑动窗口 👍 787 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	// 滑动窗口，初始化为乘法单位元
	windowProduct := 1
	// 记录符合条件的子数组个数
	count := 0

	for right < len(nums) {
		// 扩大窗口，并更新窗口数据
		windowProduct *= nums[right]
		right++

		//注意这个条件，不满足时停止循环，使缩小窗口停止的条件
		for left < right && windowProduct >= k {
			// 缩小窗口，并更新窗口数据
			windowProduct /= nums[left]
			left++
		}
		// 现在必然是一个合法的窗口，但注意思考这个窗口中的子数组个数怎么计算：
		// 比方说 left = 1, right = 4 划定了 [1, 2, 3] 这个窗口（right 是开区间）
		// 但不止 [left..right] 是合法的子数组，[left+1..right], [left+2..right] 等都是合法子数组
		// 所以我们需要把 [3], [2,3], [1,2,3] 这 right - left 个子数组都加上
		count += right - left
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubarrayProductLessThanK(t *testing.T) {

}
