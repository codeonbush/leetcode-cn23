package leetcode

import (
	"testing"
)

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。






 示例 1：


输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

 示例 2：


输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]




 提示：


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums 已按 非递减顺序 排序




 进阶：


 请你设计时间复杂度为 O(n) 的算法解决本问题


 Related Topics 数组 双指针 排序 👍 1000 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 平方的特点是会把负数变成正数，所以一个负数和一个正数平方后的大小要根据绝对值来比较。
//
// 可以把元素 0 作为分界线，0 左侧的负数是一个有序数组 nums1，
// 0 右侧的正数是另一个有序数组 nums2，那么这道题就和
// 88. 合并两个有序数组 差不多，
//
// 所以，我们可以去寻找正负数的分界点，然后向左右扩展，执行合并有序数组的逻辑。
// 不过还有个更好的办法，不用找正负分界点，
// 而是直接将双指针分别初始化在 nums 的开头和结尾，相当于合并两个从大到小排序的数组，
func sortedSquares(nums []int) []int {
	n := len(nums)
	// 两个指针分别初始化在正负子数组绝对值最大的元素索引
	i, j := 0, n-1
	// 得到的有序结果是降序的，两个降序列表
	p := n - 1
	res := make([]int, n)
	// 执行双指针合并有序数组的逻辑
	for i <= j {
		if abs(nums[i]) > abs(nums[j]) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
		p--
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSquaresOfASortedArray(t *testing.T) {

}
