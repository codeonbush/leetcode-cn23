package leetcode

import (
	"testing"
)

/**
给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。

 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。

 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。



 示例 1：


输入：n = 7
输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0
],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]


 示例 2：


输入：n = 3
输出：[[0,0,0]]




 提示：


 1 <= n <= 20


 Related Topics 树 递归 记忆化搜索 动态规划 二叉树 👍 406 👎 0

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
//隐含条件，1.所有节点的val=0
//构建n个节点的二叉树=从1个节点到n个节点构建二叉树
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		// 题目描述的满二叉树不可能是偶数个节点
		return []*TreeNode{}
	}
	memo := make([][]*TreeNode, n+1)
	return buildT(n, memo)
}

// 定义：输入一个 n，生成节点数为 n 的所有可能的满二叉树
func buildT(n int, memo [][]*TreeNode) []*TreeNode {
	res := []*TreeNode{}
	// base case
	if n == 1 {
		res = append(res, &TreeNode{Val: 0})
		return res
	}
	if memo[n] != nil {
		// 避免冗余计算
		return memo[n]
	}

	// 递归生成所有符合条件的左右子树
	for i := 1; i < n; i += 2 {
		//i是左子树的节点数量，j为右子树节点数量
		j := n - i - 1
		// 利用函数定义，生成左右子树
		leftSubTrees := buildT(i, memo)
		rightSubTrees := buildT(j, memo)
		// 左右子树的不同排列也能构成不同的二叉树
		for _, left := range leftSubTrees {
			for _, right := range rightSubTrees {
				// 生成根节点
				root := &TreeNode{Val: 0}
				// 组装出一种可能的二叉树形状
				root.Left = left
				root.Right = right
				// 加入结果列表
				res = append(res, root)
			}
		}
	}
	// 存入备忘录
	memo[n] = res
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAllPossibleFullBinaryTrees(t *testing.T) {

}
