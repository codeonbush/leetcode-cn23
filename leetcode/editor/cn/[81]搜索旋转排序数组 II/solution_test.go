package leetcode

import (
	"testing"
)

/**
已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。

 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[
k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,
4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

 给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值
target ，则返回 true ，否则返回 false 。

 你必须尽可能减少整个操作步骤。



 示例 1：


输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true


 示例 2：


输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false



 提示：


 1 <= nums.length <= 5000
 -10⁴ <= nums[i] <= 10⁴
 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 -10⁴ <= target <= 10⁴




 进阶：


 此题与 搜索旋转排序数组 相似，但本题中的 nums 可能包含 重复 元素。这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？




 Related Topics 数组 二分查找 👍 781 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		// 本题需要在计算 mid 之前收缩左右边界去重
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		// 其余逻辑和第 33 题完全相同
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
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
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArrayIi(t *testing.T) {

}
