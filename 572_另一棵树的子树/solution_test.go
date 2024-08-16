package leetcode

import (
	"testing"
)

/**


 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返
回 false 。



 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。



 示例 1：


输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true


 示例 2：


输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false




 提示：


 root 树上的节点数量范围是 [1, 2000]
 subRoot 树上的节点数量范围是 [1, 1000]
 -10⁴ <= root.val <= 10⁴
 -10⁴ <= subRoot.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 1030 👎 0

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
//func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//	res := false
//	//检查以root的为根节点的树和subroot是否相同
//	var traverse func(root *TreeNode)
//	traverse = func(root *TreeNode) {
//		if root == nil && subRoot != nil {
//			return
//		}
//		if isSame(root, subRoot) {
//			res = true
//			return
//		}
//		traverse(root.Left)
//		traverse(root.Right)
//	}
//	traverse(root)
//	return res
//}
//
//func isSame(root *TreeNode, subRoot *TreeNode) bool {
//	if root == nil && subRoot == nil {
//		return true
//	}
//	if root == nil || subRoot == nil {
//		return false
//	}
//	if root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right) {
//		return true
//	}
//	return false
//}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	// 判断以 root 为根的二叉树是否和 subRoot 相同
	if isSameTree(root, subRoot) {
		return true
	}
	// 去左右子树中判断是否有和 subRoot 相同的子树
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 判断一对节点是否相同
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// 判断其他节点是否相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubtreeOfAnotherTree(t *testing.T) {

}
