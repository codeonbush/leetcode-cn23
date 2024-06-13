package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
nums[i] 的元素的数量。



 示例 1：


输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素


 示例 2：


输入：nums = [-1]
输出：[0]


 示例 3：


输入：nums = [-1,-1]
输出：[0,0]




 提示：


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴


 Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1079 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
//

//func countSmaller(nums []int) []int {
//	count := make([]int, 0, len(nums))
//	result := make([]int, len(nums))
//	for i := range result {
//		result[i] = i
//	}
//	sort(result, count)
//	return count
//}
//
//func sort(nums, count []int) []int {
//	if len(nums) == 0 {
//		return nums
//	}
//	mid := len(nums) / 2
//	left := sort(nums[:mid], count)
//	right := sort(nums[mid:], count)
//	return merge(left, right, count)
//}
//
//func merge(a, b, count []int) []int {
//	res := make([]int, 0, len(a)+len(b))
//	i, j, p := 0, 0, 0
//	for i < len(a) || j < len(b) {
//		if a[i] < b[j] {
//			res[p] = a[i]
//			count[a[i]] += j
//			i++
//		} else {
//			res[p] = b[j]
//			j++
//		}
//		p++
//	}
//	res = append(res, a[i:]...)
//	res = append(res, b[j:]...)
//	return res
//}

// - 使用 `indices` 数组来追踪原始数组的索引，避免直接操作原始数组。
// - 在 `merge` 函数中，正确地更新 `counts` 数组，确保每个元素右侧比它小的元素个数被正确计算。
// - 避免了数组越界的问题，确保合并过程中的索引操作安全。
func countSmaller(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)
	//初始化 `counts` 数组用于存储结果。
	//   - 初始化 `indices` 数组用于存储原始数组的索引。
	//   - 调用 `mergeSort` 函数进行排序和计数。
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	mergeSort(nums, indices, counts)
	return counts
}

func mergeSort(nums, indices, counts []int) []int {
	if len(indices) <= 1 {
		return indices
	}
	mid := len(indices) / 2
	left := mergeSort(nums, indices[:mid], counts)
	right := mergeSort(nums, indices[mid:], counts)
	return merge(nums, left, right, counts)
}

func merge(nums []int, left, right, counts []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if nums[left[i]] <= nums[right[j]] {
			counts[left[i]] += j
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	for i < len(left) {
		counts[left[i]] += j
		merged = append(merged, left[i])
		i++
	}
	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}
	return merged
}

//func countSmaller(nums []int) []int {
//	n := len(nums)
//	counts := make([]int, n)
//	var merge func(left, right []int) []int
//	merge = func(left, right []int) []int {
//		i, j := 0, 0
//		merged := make([]int, 0, len(left)+len(right))
//		for i < len(left) || j < len(right) {
//			//left和right都是索引切片的一部分，索引切片中的【值】代表元素在原始切片nums中的位置
//			if i < len(left) && (j == len(right) || nums[left[i]] <= nums[right[j]]) {
//				counts[left[i]] += j
//				merged = append(merged, left[i])
//				i++
//			} else {
//				merged = append(merged, right[j])
//				j++
//			}
//		}
//		return merged
//	}
//	var mergeSort func([]int) []int
//	mergeSort = func(arr []int) []int {
//		size := len(arr)
//		if size <= 1 {
//			return arr
//		}
//		mid := size / 2
//		left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
//		return merge(left, right)
//	}
//
//	mergeSort(sliceIndex(n))
//	return counts
//}
//
//func sliceIndex(length int) []int {
//	result := make([]int, length)
//	for i := range result {
//		result[i] = i
//	}
//	return result
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountOfSmallerNumbersAfterSelf(t *testing.T) {

}
