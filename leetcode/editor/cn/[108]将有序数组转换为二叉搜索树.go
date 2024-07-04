package leetcode

import (
	"testing"
)

/**
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。



 示例 1：


输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：



 示例 2：


输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。




 提示：


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums 按 严格递增 顺序排列


 Related Topics 树 二叉搜索树 数组 分治 二叉树 👍 1528 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return buildT(nums, 0, len(nums)-1)
}

// 注意使用索引和传递数组的区别，索引前后都是闭区间，go切片前闭后开
// 所以使用索引方案更适配多种语言
func buildT(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	rootIndex := (end + start) / 2
	root := &TreeNode{Val: nums[rootIndex]}
	root.Left = buildT(nums, start, rootIndex-1)
	root.Right = buildT(nums, rootIndex+1, end)
	return root
}

//func sortedArrayToBST(nums []int) *TreeNode {
//	return build(nums, 0, len(nums)-1)
//}
//
//// 将闭区间 [left, right] 中的元素转化成 BST，返回根节点
//func build(nums []int, left int, right int) *TreeNode {
//	if left > right {
//		// 区间为空
//		return nil
//	}
//	// 构造根节点
//	// BST 节点左小右大，中间的元素就是根节点
//	mid := (left + right) / 2
//	root := &TreeNode{Val: nums[mid]}
//	// 递归构建左子树
//	root.Left = build(nums, left, mid-1)
//	// 递归构造右子树
//	root.Right = build(nums, mid+1, right)
//
//	return root
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertSortedArrayToBinarySearchTree(t *testing.T) {

}
