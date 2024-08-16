package leetcode

import (
	"testing"
)

/**
给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。



 示例 1：


输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

 示例 2：


输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。



 提示：


 1 <= nums.length <= 10⁵
 nums[i] 不是 0 就是 1
 0 <= k <= nums.length


 Related Topics 数组 二分查找 前缀和 滑动窗口 👍 712 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	// 记录窗口中 1 的出现次数
	windowOneCount := 0
	// 记录结果长度
	res := 0

	// 开始滑动窗口模板
	for right < len(nums) {
		// 扩大窗口
		if nums[right] == 1 {
			windowOneCount++
		}
		right++

		for right-left-windowOneCount > k {
			// 当窗口中需要替换的 0 的数量大于 k，缩小窗口
			if nums[left] == 1 {
				windowOneCount--
			}
			left++
		}
		// 此时一定是一个合法的窗口，求最大窗口长度
		res = max(res, right-left)
	}
	return res
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaxConsecutiveOnesIii(t *testing.T) {

}
