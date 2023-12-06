package leetcode

import (
	"math"
	"testing"
)

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
// 示例 2：
//
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
// 示例 3：
//
// 输入：root = []
// 输出：true
//
// 提示：
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 二叉树 👍 1457 👎 0
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归方案
//一定要体现【子树也要满足】这个条件，这是递归的本质，满足这个条件就是在用隐藏栈遍历子树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//一定要体现【子树也要满足】这个条件
	//当前节点要做的事：比较左右子树的高度
	//对于子节点，也要判断是否平衡
	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBalancedBinaryTree(t *testing.T) {
	//cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d expected %d, but %d got",
	//				c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}
}
