package leetcode

import (
	"testing"
)

/**
整数数组 nums 按升序排列，数组中的值 互不相同 。

 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+
1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4
,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



 示例 1：


输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4


 示例 2：


输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

 示例 3：


输入：nums = [1], target = 0
输出：-1




 提示：


 1 <= nums.length <= 5000
 -10⁴ <= nums[i] <= 10⁴
 nums 中的每个值都 独一无二
 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 -10⁴ <= target <= 10⁴


 Related Topics 数组 二分查找 👍 2978 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	// 左右都闭的搜索区间
	left, right := 0, len(nums)-1
	// 因为是闭区间，所以结束条件为 left > right
	for left <= right {
		mid := left + (right-left)/2
		// 首先检查 nums[mid]，是否找到 target
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			// mid 落在断崖左边，此时 nums[left..mid] 有序
			if target >= nums[left] && target < nums[mid] {
				// target 落在 [left..mid-1] 中
				right = mid - 1
			} else {
				// target 落在 [mid+1..right] 中
				left = mid + 1
			}
		} else {
			// mid 落在断崖右边，此时 nums[mid..right] 有序
			if target <= nums[right] && target > nums[mid] {
				// target 落在 [mid+1..right] 中
				left = mid + 1
			} else {
				// target 落在 [left..mid-1] 中
				right = mid - 1
			}
		}
	}
	// while 结束还没找到，说明 target 不存在
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArray(t *testing.T) {

}
