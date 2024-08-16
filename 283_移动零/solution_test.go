package leetcode

import (
	"testing"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
//输入: nums = [0]
//输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
//
// Related Topics 数组 双指针 👍 2340 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	lenN := len(nums)
	if lenN == 0 {
		return
	}
	count := removeElem(nums, 0)
	for i := count; i < lenN; i++ {
		nums[i] = 0
	}
	return
}

func removeElem(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMoveZeroes(t *testing.T) {

}
