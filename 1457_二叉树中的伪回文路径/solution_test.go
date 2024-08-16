package leetcode

import (
	"testing"
)

/**
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。

 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。



 示例 1：




输入：root = [2,3,1,3,1,null,1]
输出：2
解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排列
[1,2,1] 。


 示例 2：




输入：root = [2,1,1,1,3,null,null,null,null,null,1]
输出：1
解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。


 示例 3：


输入：root = [9]
输出：1




 提示：


 给定二叉树的节点数目在范围 [1, 10⁵] 内
 1 <= Node.val <= 9


 Related Topics 位运算 树 深度优先搜索 广度优先搜索 二叉树 👍 129 👎 0

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
//这道题需要用到「遍历」的思维。
//
//遍历一遍二叉树就能得到每条路径上的数字，但这题的考点在于，如何判断一组数字是否存在一个回文串组合？
//
//稍加思考不难想到：如果一组数字中，只有最多一个数字出现的次数为奇数，剩余数字的出现次数均为偶数，那么这组数字可以组成一个回文串。
//
//题目说了 1 <= root.val <= 9，所以我们可以用一个大小为 10 的 count 数组做计数器来记录每条路径上的元素出现次数，到达叶子节点之后根据元素出现的次数判断是否可以构成回文串。
//
//当然，我们也可以用更巧妙的位运算来实现上述逻辑：
//
//1、首先用到异或运算的特性，1 ^ 1 = 0, 0 ^ 0 = 0, 1 ^ 0 = 1。
//
//2、其次用到 n & (n - 1) 去除二进制最后一个 1 的技巧
//- `count = count ^ (1 << root.Val)`：对于每个节点 `root`，将其值 `root.Val` 左移1位后，通过异或操作更新 `count`。这一步用以跟踪每个数字的出现次数，如果数值相应的位已经为1（表示之前出现过一次），再次异或将其变回0（即成对移除，反映偶数次出现）；反之则设为1（表示独立出现，反映奇数次出现）。
//
//- 当到达一个叶子节点时（即 `root.Left == nil && root.Right == nil`），使用 `(count & (count - 1)) == 0` 这一条件检查 `count`。这一操作是检查 `count` 是否在二进制表示中最多只有一个位是1，如果是，则路径字符组合可以形成回文排列，将结果 `res` 增1。
func pseudoPalindromicPaths(root *TreeNode) int {
	res := 0
	var traverse2 func(root *TreeNode, count int)
	traverse2 = func(root *TreeNode, count int) { // 二叉树遍历函数
		if root == nil {
			return
		}
		count = count ^ (1 << root.Val) // 用位运算计数
		if root.Left == nil && root.Right == nil {
			// 判断二进制中只有一位 1，原理见 https://mp.weixin.qq.com/s/8lJNdnJ0tWm2CapiW_u7XA
			if count&(count-1) == 0 {
				res++
			}
		}
		traverse2(root.Left, count)
		traverse2(root.Right, count)
	}
	traverse2(root, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPseudoPalindromicPathsInABinaryTree(t *testing.T) {

}
