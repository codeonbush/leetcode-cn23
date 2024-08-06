package leetcode

import(
    "testing"
)

/**
给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。 

 返回峰值元素的下标。 

 你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。 

 

 示例 1： 

 
输入：arr = [0,1,0]
输出：1
 

 示例 2： 

 
输入：arr = [0,2,1,0]
输出：1
 

 示例 3： 

 
输入：arr = [0,10,5,2]
输出：1
 

 

 提示： 

 
 3 <= arr.length <= 10⁵ 
 0 <= arr[i] <= 10⁶ 
 题目数据 保证 arr 是一个山脉数组 
 

 Related Topics 数组 二分查找 👍 403 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(nums []int) int {
    // 取两端都闭的二分搜索
    left, right := 0, len(nums)-1
    // 因为题目必然有解，所以设置 left == right 为结束条件
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[mid+1] {
            // mid 本身就是峰值或其左侧有一个峰值
            right = mid
        } else {
            // mid 右侧有一个峰值
            left = mid + 1
        }
    }
    return left
}
//leetcode submit region end(Prohibit modification and deletion)


func TestPeakIndexInAMountainArray(t *testing.T){
    
}
