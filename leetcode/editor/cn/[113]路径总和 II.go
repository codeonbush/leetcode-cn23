package leetcode

import (
	"testing"
)

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

 叶子节点 是指没有子节点的节点。







 示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]


 示例 2：


输入：root = [1,2,3], targetSum = 5
输出：[]


 示例 3：


输入：root = [1,2], targetSum = 0
输出：[]




 提示：


 树中节点总数在范围 [0, 5000] 内
 -1000 <= Node.val <= 1000
 -1000 <= targetSum <= 1000


 Related Topics 树 深度优先搜索 回溯 二叉树 👍 1110 👎 0

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
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	var traverse func(root *TreeNode, path []int, pathSum int)
	traverse = func(root *TreeNode, path []int, pathSum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		pathSum += root.Val
		if root.Left == nil && root.Right == nil {
			if pathSum == targetSum {
				//res = append(res, path)
				//拷贝一份以避免随后的修改
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
		}
		traverse(root.Left, path, pathSum)
		traverse(root.Right, path, pathSum)
		path = path[:len(path)-1]
		pathSum = pathSum - root.Val
	}
	traverse(root, path, 0)
	return res
}

//func pathSum(root *TreeNode, targetSum int) [][]int {
//	var res [][]int
//	var traverse func(*TreeNode, int, []int)
//
//	traverse = func(node *TreeNode, currSum int, path []int) {
//		if node == nil {
//			return
//		}
//
//		// 追加当前节点值到路径
//		currSum += node.Val
//		path = append(path, node.Val)
//
//		if node.Left == nil && node.Right == nil {
//			// 只有在叶子节点我们才检查是否路径和匹配
//			if currSum == targetSum {
//				// 拷贝一份以避免随后的修改
//				tmp := make([]int, len(path))
//				copy(tmp, path)
//				res = append(res, tmp)
//			}
//			return
//		}
//
//		traverse(node.Left, currSum, path)
//		traverse(node.Right, currSum, path)
//	}
//
//	traverse(root, 0, []int{})
//	return res
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSumIi(t *testing.T) {

}
