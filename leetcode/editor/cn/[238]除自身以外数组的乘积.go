package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。

 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



 示例 1:


输入: nums = [1,2,3,4]
输出: [24,12,8,6]


 示例 2:


输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]




 提示：


 2 <= nums.length <= 10⁵
 -30 <= nums[i] <= 30
 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内




 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

 Related Topics 数组 前缀和 👍 1811 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
//优化后的代码可以在空间复杂度上进行更高效的处理，
//不必使用额外的前缀积和后缀积数组，只使用一个结果数组来存储中间计算結果。
// 这种方法通过两遍遍历解决问题，第一遍遍历用于计算每个元素的前缀积，
// 第二遍遍历用于计算每个元素的后缀积并同时将其与之前计算的前缀积相乘，
// 从而得到最终结果。这减少了额外的空间使用，使空间复杂度降为 O(1)（不包括返回的结果数组）。
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 从左到右计算前缀积
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 从右到左计算后缀积并且和前缀积相乘
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}

//func productExceptSelf(nums []int) []int {
//	n := len(nums)
//	// 从左到右的前缀积，prefix[i] 是 nums[0..i] 的元素积
//	prefix := make([]int, n)
//	prefix[0] = nums[0]
//	for i := 1; i < n; i++ {
//		prefix[i] = prefix[i-1] * nums[i]
//	}
//	// 从右到左的前缀积，suffix[i] 是 nums[i..n-1] 的元素积
//	suffix := make([]int, n)
//	suffix[n-1] = nums[n-1]
//	for i := n - 2; i >= 0; i-- {
//		suffix[i] = suffix[i+1] * nums[i]
//	}
//	// 结果数组
//	res := make([]int, n)
//	res[0] = suffix[1]
//	res[n-1] = prefix[n-2]
//	for i := 1; i < n-1; i++ {
//		// 除了 nums[i] 自己的元素积就是 nums[i] 左侧和右侧所有元素之积
//		res[i] = prefix[i-1] * suffix[i+1]
//	}
//	return res
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestProductOfArrayExceptSelf(t *testing.T) {

}
