package leetcode

import "testing"

/**
给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。

 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。



 示例 1：



 输入：root = [1,2,3,4,5,6]
输出：110
解释：删除红色的边，得到 2 棵子树，和分别为 11 和 10 。它们的乘积是 110 （11*10）


 示例 2：



 输入：root = [1,null,2,3,4,null,null,5,6]
输出：90
解释：移除红色的边，得到 2 棵子树，和分别是 15 和 6 。它们的乘积为 90 （15*6）


 示例 3：

 输入：root = [2,3,9,10,7,8,6,5,4,11,1]
输出：1025


 示例 4：ff

 输入：root = [1,1]
输出：1




 提示：


 每棵树最多有 50000 个节点，且至少有 2 个节点。
 每个节点的值在 [1, 10000] 之间。


 Related Topics 树 深度优先搜索 二叉树 👍 103 👎 0

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
//在二叉树中切出一个小二叉树（子树），计算这个子树节点之和与剩下的节点之和的乘积。
//想求最大乘积，那就穷举，把所有可能的切法都穷举一遍，计算乘积。
//
//任何子树的节点之和都可以在后序位置获得，而剩下的其他节点之和就是整棵二叉树的节点之和减去子树节点之和。
//
//所以我们写一个 getSum 函数，先执行一遍计算整棵树的节点之和，然后再调用一次利用它的后序位置计算乘积。
//关键在于，需要先获取整棵树的节点之和

//func maxProduct(root *TreeNode) int {
//	res := int64(0)
//	treeSum := getSum(root, &res, 0)
//	getSum(root, &res, treeSum)
//	return int(res % (1e9 + 7))
//}
//
//func getSum(root *TreeNode, res *int64, treeSum int) int {
//	if root == nil {
//		return 0
//	}
//
//	leftSum := getSum(root.Left, res, treeSum)
//	rightSum := getSum(root.Right, res, treeSum)
//	rootSum := leftSum + rightSum + root.Val
//
//	// 后序位置计算乘积
//	*res = max2(*res, int64(rootSum)*(int64(treeSum)-int64(rootSum)))
//	return rootSum
//}
//
//func max2(a, b int64) int64 {
//	if a > b {
//		return a
//	}
//	return b
//}

func maxProduct(root *TreeNode) int {
	const mod int = 1e9 + 7
	var totalSum int

	// 第一次递归计算整棵树的节点之和
	var sumTree func(node *TreeNode) int
	sumTree = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := sumTree(node.Left)
		rightSum := sumTree(node.Right)
		totalSum += node.Val
		return leftSum + rightSum + node.Val
	}

	totalSum = sumTree(root)

	maxProd := int64(0)
	// 第二次递归计算每个子树的节点之和，比较乘积并获取最大值
	var calculateAndCompare func(node *TreeNode) int
	calculateAndCompare = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := calculateAndCompare(node.Left)
		rightSum := calculateAndCompare(node.Right)
		subtreeSum := leftSum + rightSum + node.Val
		// 计算这个子树切割后的乘积
		maxProd = max2(maxProd, int64(subtreeSum)*(int64(totalSum)-int64(subtreeSum)))
		return subtreeSum
	}

	calculateAndCompare(root)
	return int(maxProd % int64(mod))
}

func max2(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumProductOfSplittedBinaryTree(t *testing.T) {

}
