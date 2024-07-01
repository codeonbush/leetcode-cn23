package leetcode

import (
	"testing"
)

/**
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。


 更正式地说，即 root.val = min(root.left.val, root.right.val) 总成立。

 给出这样的一个二叉树，你需要输出所有节点中的 第二小的值 。

 如果第二小的值不存在的话，输出 -1 。



 示例 1：


输入：root = [2,2,5,null,null,5,7]
输出：5
解释：最小的值是 2 ，第二小的值是 5 。


 示例 2：


输入：root = [2,2,2]
输出：-1
解释：最小的值是 2, 但是不存在第二小的值。




 提示：


 树中节点数目在范围 [1, 25] 内
 1 <= Node.val <= 2³¹ - 1
 对于树中每个节点 root.val == min(root.left.val, root.right.val)


 Related Topics 树 深度优先搜索 二叉树 👍 308 👎 0

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
//这道题需要用到「分解问题」的思维。
//
//这题很有意思，按照这棵二叉树的特性，根节点无疑是最小的那个元素，但问题是第二小的那个元素在哪里呢？
//
//如果根节点的左右子节点的值不同，根节点的值就是较小的那个节点（假设是左子节点）的值，那么较大的那个节点（右子节点）的值是不是就一定是第二小的那个元素？
//
//不一定，第二小的值也可能在左子树中，这个节点是左子树中第二小的节点：
//类似的道理，如果根节点的左右子节点相同，那么需要把左右子树的第二小元素计算出来，比较其中较小的元素，作为整棵树的第二小元素：
// 定义：输入一棵二叉树，返回这棵二叉树中第二小的节点值
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return -1
	}
	// 左右子节点中不同于 root.val 的那个值可能是第二小的值
	left, right := root.Left.Val, root.Right.Val
	// 如果左右子节点都等于 root.val，则去左右子树递归寻找第二小的值
	if root.Val == root.Left.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == root.Right.Val {
		right = findSecondMinimumValue(root.Right)
	}
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	// 如果左右子树都找到一个第二小的值，更小的那个是整棵树的第二小元素
	return min(left, right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSecondMinimumNodeInABinaryTree(t *testing.T) {

}
