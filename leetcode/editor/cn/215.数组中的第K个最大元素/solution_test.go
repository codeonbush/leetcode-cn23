package leetcode

import (
	"testing"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2369 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

/**
 * @Description l 左索引 r右索引 k目标排序
 **/
//在快速排序交换元素的过程中，切分元素满足左侧小于他，右侧大于他
func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	//初始化为左索引,可以应对大量重复元素的数组
	p := nums[l]
	//使用双指针寻找元素并交换
	i := l - 1
	j := r + 1
	//双指针相遇时，本次排序结束
	for i < j {
		for i++; nums[i] < p; i++ {
		}
		for j--; nums[j] > p; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAnArray(t *testing.T) {
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
