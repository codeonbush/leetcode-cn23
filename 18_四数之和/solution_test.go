package leetcode

import (
	"sort"
	"testing"
)

/**
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b],
 nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：


 0 <= a, b, c, d < n
 a、b、c 和 d 互不相同
 nums[a] + nums[b] + nums[c] + nums[d] == target


 你可以按 任意顺序 返回答案 。



 示例 1：


输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]


 示例 2：


输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]




 提示：


 1 <= nums.length <= 200
 -10⁹ <= nums[i] <= 10⁹
 -10⁹ <= target <= 10⁹


 Related Topics 数组 双指针 排序 👍 1927 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4, 0)
}

// 注意：调用这个函数之前一定要先给 nums 排序
// n 填写想求的是几数之和，start 从哪个索引开始计算（一般填 0），target 填想凑出的目标和
func nSum(nums []int, target, n, start int) [][]int {
	var res [][]int
	if n < 2 || len(nums) < n {
		return res
	}
	if n == 2 {
		return twoSum(nums, target, start)
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subRes := nSum(nums, target-nums[i], n-1, i+1)
		for _, sub := range subRes {
			res = append(res, append(sub, nums[i]))
		}
	}
	return res
}

func twoSum(nums []int, target, start int) [][]int {
	lo, hi := start, len(nums)-1
	var res [][]int
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target {
			lo++
		} else if sum > target {
			hi--
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
			for lo < hi && nums[hi] == nums[hi+1] {
				hi--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFourSum(t *testing.T) {

}
