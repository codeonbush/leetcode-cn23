package leetcode

import (
	"testing"
)

/**
我们可以为二叉树 T 定义一个 翻转操作 ，如下所示：选择任意节点，然后交换它的左子树和右子树。

 只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转 等价 于二叉树 Y。

 这些树由根节点 root1 和 root2 给出。如果两个二叉树是否是翻转 等价 的函数，则返回 true ，否则返回 false 。



 示例 1：




输入：root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,
null,null,null,8,7]
输出：true
解释：我们翻转值为 1，3 以及 5 的三个节点。


 示例 2:


输入: root1 = [], root2 = []
输出: true


 示例 3:


输入: root1 = [], root2 = [1]
输出: false




 提示：


 每棵树节点数在 [0, 100] 范围内
 每棵树中的每个值都是唯一的、在 [0, 99] 范围内的整数


 Related Topics 树 深度优先搜索 二叉树 👍 161 👎 0

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

// 翻转的意思就是a的左子树=b的右子树
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	// 定义：输入两棵二叉树，判断这两棵二叉树是否是翻转等价的
	// 判断 root1 和 root2 两个节点是否能够匹配
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	// 根据函数定义，判断子树是否能够匹配
	// 不翻转、翻转两种情况满足一种即可算是匹配
	return func() bool {
		// 不翻转子树
		if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
			return true
		}
		// 反转子树
		if flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) {
			return true
		}
		return false
	}()
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFlipEquivalentBinaryTrees(t *testing.T) {

}
