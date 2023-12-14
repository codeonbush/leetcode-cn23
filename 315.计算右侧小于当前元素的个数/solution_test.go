package leetcode

import (
	"testing"
)

//给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
//nums[i] 的元素的数量。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,2,6,1]
//输出：[2,1,1,0]
//解释：
//5 的右侧有 2 个更小的元素 (2 和 1)
//2 的右侧仅有 1 个更小的元素 (1)
//6 的右侧有 1 个更小的元素 (1)
//1 的右侧有 0 个更小的元素
//
//
// 示例 2：
//
//
//输入：nums = [-1]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：nums = [-1,-1]
//输出：[0,0]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1035 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

//测试暴力解法超时
//func countSmaller(nums []int) []int {
//	result := make([]int, len(nums), len(nums))
//	if len(nums) == 0 {
//		return result
//	}
//	for i := 0; i < len(nums); i++ {
//		count := 0
//		for j := i; j < len(nums); j++ {
//			if nums[j] < nums[i] {
//				count++
//			}
//		}
//		result[i] = count
//	}
//	return result
//}

// 逆序对数量：
func countSmaller(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)
	var merge func(left, right []int) []int
	merge = func(left, right []int) []int {
		i, j := 0, 0
		merged := make([]int, 0, len(left)+len(right))
		for i < len(left) || j < len(right) {
			//left和right都是索引切片的一部分，索引切片中的【值】代表元素在原始切片nums中的位置
			if i < len(left) && (j == len(right) || nums[left[i]] <= nums[right[j]]) {
				counts[left[i]] += j
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		}
		return merged
	}
	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		size := len(arr)
		if size <= 1 {
			return arr
		}
		mid := size / 2
		left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
		return merge(left, right)
	}

	mergeSort(sliceIndex(n))
	return counts
}

func sliceIndex(length int) []int {
	result := make([]int, length)
	for i := range result {
		result[i] = i
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountOfSmallerNumbersAfterSelf(t *testing.T) {
	//testnums := []int{5, 4, 3, 2, 1}
}
