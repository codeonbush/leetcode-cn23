package leetcode

import (
	"testing"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



 示例 1:


输入: [3,2,1,5,6,4], k = 2
输出: 5


 示例 2:


输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4



 提示：


 1 <= k <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴


 Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2500 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 1. 使用三向切分（Dutch National Flag partitioning）来处理重复元素。
// 2. `partition` 函数返回两个索引 `lt` 和 `gt`，分别表示小于和大于 `pivot` 的边界。
// 3. 在 `quickSelect` 中，根据 `k` 的位置选择递归的区间，减少不必要的递归调用。
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}

	lt, gt := partition(nums, l, r)

	if k >= lt && k <= gt {
		return nums[k]
	} else if k < lt {
		return quickSelect(nums, l, lt-1, k)
	} else {
		return quickSelect(nums, gt+1, r, k)
	}
}

func partition(nums []int, l, r int) (int, int) {
	pivot := nums[l]
	lt, i, gt := l, l+1, r

	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAnArray(t *testing.T) {

}
