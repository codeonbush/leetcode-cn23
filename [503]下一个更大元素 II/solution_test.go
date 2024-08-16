package leetcode

import (
	"testing"
)

/**
给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。


 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。




 示例 1:


输入: nums = [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。


 示例 2:


输入: nums = [1,2,3,4,3]
输出: [2,3,4,-1,4]




 提示:


 1 <= nums.length <= 10⁴
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 栈 数组 单调栈 👍 995 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	// 用数组模拟栈
	s := make([]int, 0)

	// 数组长度加倍模拟环形数组
	for i := 2*n - 1; i >= 0; i-- {
		// 索引 i 要求模，其他的和模板一样
		for len(s) > 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1] // pop element from stack
		}

		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1]
		}

		s = append(s, nums[i%n]) // push element to stack
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterElementIi(t *testing.T) {

}
