package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree with unique values and the values of two
different nodes of the tree x and y, return true if the nodes corresponding to the
values x and y in the tree are cousins, or false otherwise.

 Two nodes of a binary tree are cousins if they have the same depth with
different parents.

 Note that in a binary tree, the root node is at the depth 0, and children of
each depth k node are at the depth k + 1.


 Example 1:


Input: root = [1,2,3,4], x = 4, y = 3
Output: false


 Example 2:


Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true


 Example 3:


Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false



 Constraints:


 The number of nodes in the tree is in the range [2, 100].
 1 <= Node.val <= 100
 Each node has a unique value.
 x != y
 x and y are exist in the tree.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 344 👎 0

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
func isCousins(root *TreeNode, x int, y int) bool {
	//深度相同，且父节点不同
	xDepth := 0
	yDepth := 0
	xParent, yParent := new(TreeNode), new(TreeNode)
	var traverse func(root *TreeNode, depth int, parent *TreeNode)
	traverse = func(root *TreeNode, depth int, parent *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == x {
			xDepth = depth
			xParent = parent
		}
		if root.Val == y {
			yDepth = depth
			yParent = parent
		}
		traverse(root.Left, depth+1, root)
		traverse(root.Right, depth+1, root)
	}
	traverse(root, 0, nil)
	if xDepth == yDepth && xParent != yParent {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCousinsInBinaryTree(t *testing.T) {

}
