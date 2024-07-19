package leetcode

import (
	"math"
	"testing"
)

/**
给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。

 找出满足下述条件的下标对 (i, j)：


 i != j,
 abs(i - j) <= indexDiff
 abs(nums[i] - nums[j]) <= valueDiff


 如果存在，返回 true ；否则，返回 false 。



 示例 1：


输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
输出：true
解释：可以找出 (i, j) = (0, 3) 。
满足下述 3 个条件：
i != j --> 0 != 3
abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0


 示例 2：


输入：nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
输出：false
解释：尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。




 提示：


 2 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 1 <= indexDiff <= nums.length
 0 <= valueDiff <= 10⁹


 Related Topics 数组 桶排序 有序集合 排序 滑动窗口 👍 734 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}

	getID := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1
	}

	window := make(map[int]int)
	w := t + 1

	for i := 0; i < len(nums); i++ {
		m := getID(nums[i], w)

		// 为了防止 i == j，所以在扩大窗口之前先判断是否有符合题意的索引对 (i, j)
		// 查找略大于 nums[right] 的那个元素
		if _, ok := window[m]; ok {
			return true
		}
		// 查找略小于 nums[right] 的那个元素
		if v, ok := window[m-1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}
		if v, ok := window[m+1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}

		// 扩大窗口
		window[m] = nums[i]

		if i >= k {
			// 缩小窗口
			delete(window, getID(nums[i-k], w))
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainsDuplicateIii(t *testing.T) {

}
