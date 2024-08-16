package leetcode

import (
	"testing"
)

/**
Given an integer array nums, handle multiple queries of the following type:


 Calculate the sum of the elements of nums between indices left and right
inclusive where left <= right.


 Implement the NumArray class:


 NumArray(int[] nums) Initializes the object with the integer array nums.
 int sumRange(int left, int right) Returns the sum of the elements of nums
between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... +
nums[right]).



 Example 1:


Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁵ <= nums[i] <= 10⁵
 0 <= left <= right < nums.length
 At most 10⁴ calls will be made to sumRange.


 Related Topics 设计 数组 前缀和 👍 632 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	// 前缀和数组
	preSum []int
}

/* 输入一个数组，构造前缀和 */
//func Constructor(nums []int) NumArray {
//	// preSum[0] = 0，便于计算累加和
//	//前缀和不计算索引i元素本身
//	n := len(nums)
//	preSum := make([]int, n+1)
//	// 计算 nums 的累加和
//	for i := 1; i < n+1; i++ {
//		//上一个前缀和加上前一个元素
//		preSum[i] = preSum[i-1] + nums[i-1]
//	}
//
//	return NumArray{preSum: preSum}
//}

/* 查询闭区间 [left, right] 的累加和 */
func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRangeSumQueryImmutable(t *testing.T) {

}
