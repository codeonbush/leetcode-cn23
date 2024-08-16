package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2
[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。



 示例 1：


输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

 示例 2：


输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。




 提示：


 1 <= nums1.length <= nums2.length <= 1000
 0 <= nums1[i], nums2[i] <= 10⁴
 nums1和nums2中所有整数 互不相同
 nums1 中的所有整数同样出现在 nums2 中




 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？

 Related Topics 栈 数组 哈希表 单调栈 👍 1196 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 记录 nums2 中每个元素的下一个更大元素
	greater := calculateGreaterElement(nums2)
	// 转化成映射：元素 x -> x 的下一个最大元素
	greaterMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		greaterMap[nums2[i]] = greater[i]
	}
	// nums1 是 nums2 的子集，所以根据 greaterMap 可以得到结果
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = greaterMap[nums1[i]]
	}
	return res
}

func calculateGreaterElement(nums []int) []int {
	n := len(nums)
	// 存放答案的数组
	res := make([]int, n)
	s := make([]int, 0)
	// 倒着往栈里放
	for i := n - 1; i >= 0; i-- {
		// 判定个子高矮
		for len(s) != 0 && s[len(s)-1] <= nums[i] {
			// 矮个起开，反正也被挡着了。。。
			s = s[:len(s)-1]
		}
		// nums[i] 身后的更大元素
		if len(s) == 0 {
			res[i] = -1
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterElementI(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		//{
		//	nums1: []int{4, 1, 2},
		//	nums2: []int{1, 3, 4, 2},
		//	want:  []int{-1, 3, -1},
		//},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		//{
		//	nums1: []int{10, 3, 5},
		//	nums2: []int{5, 10, 3, 20, 21},
		//	want:  []int{20, 20, 10},
		//},
		//{
		//	nums1: []int{1, 2, 3, 4},
		//	nums2: []int{1, 2, 3},
		//	want:  []int{2, 3, -1, -1},
		//},
		//{
		//	nums1: []int{},
		//	nums2: []int{1, 2, 3},
		//	want:  []int{},
		//},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums1: %v, nums2: %v", tt.nums1, tt.nums2), func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
