package leetcode

import(
    "testing"
)

/**
给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。 

 整数 a 比整数 b 更接近 x 需要满足： 

 
 |a - x| < |b - x| 或者 
 |a - x| == |b - x| 且 a < b 
 

 

 示例 1： 

 
输入：arr = [1,2,3,4,5], k = 4, x = 3
输出：[1,2,3,4]
 

 示例 2： 

 
输入：arr = [1,2,3,4,5], k = 4, x = -1
输出：[1,2,3,4]
 

 

 提示： 

 
 1 <= k <= arr.length 
 1 <= arr.length <= 10⁴
 
 arr 按 升序 排列 
 -10⁴ <= arr[i], x <= 10⁴ 
 

 Related Topics 数组 双指针 二分查找 排序 滑动窗口 堆（优先队列） 👍 564 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func findClosestElements(arr []int, k int, x int) []int {
    // 二分搜索找到 x 的位置
    p := leftBound(arr, x)
    // 两端都开的区间 (left, right)
    left, right := p-1, p
    res := make([]int, 0, k)
    // 扩展区间，直到区间内包含 k 个元素
    for right-left-1 < k {
        if left == -1 {
            res = append(res, arr[right])
            right++
        } else if right == len(arr) {
            res = append([]int{arr[left]}, res...)
            left--
        } else if x-arr[left] > arr[right]-x {
            res = append(res, arr[right])
            right++
        } else {
            res = append([]int{arr[left]}, res...)
            left--
        }
    }
    return res
}

// 搜索左侧边界的二分搜索
func leftBound(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            right = mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

//leetcode submit region end(Prohibit modification and deletion)


func TestFindKClosestElements(t *testing.T){
    
}
