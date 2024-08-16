package leetcode

import (
	"testing"
)

//给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴
//
//
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 948 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	p := nums[0]
	l, r := 1, len(nums)-1
	for l <= r {
		for l <= len(nums)-1 && nums[l] < p {
			l++
		}
		for r >= 1 && nums[r] > p {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[0], nums[r] = nums[r], nums[0]
	sortArray(nums[:r])
	sortArray(nums[r+1:])
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSortAnArray(t *testing.T) {
	//cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d expected %d, but %d got",
	//				c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}
}
