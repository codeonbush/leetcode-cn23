package leetcode

import (
	"testing"
)

/**
给出二叉树的根节点 root，树上每个节点都有一个不同的值。

 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。

 返回森林中的每棵树。你可以按任意顺序组织答案。



 示例 1：




输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
输出：[[1,2,null,4],[6],[7]]


 示例 2：


输入：root = [1,2,4,null,3], to_delete = [3]
输出：[[1,2,4]]




 提示：


 树中的节点数最大为 1000。
 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
 to_delete.length <= 1000
 to_delete 包含一些从 1 到 1000、各不相同的值。


 Related Topics 树 深度优先搜索 数组 哈希表 二叉树 👍 329 👎 0

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
//首先，如果在递归过程中修改二叉树结构，必须要让父节点接收递归函数的返回值，因为子树不管删成啥样，都要接到父节点上。
//而且，可以通过函数参数传递父节点传递的数据，所以可以在前序位置判断是否得到了一个新的根节点。
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delSet := make(map[int]bool)
	for _, d := range to_delete {
		delSet[d] = true
	}
	res := make([]*TreeNode, 0)
	doDelete(root, false, &res, delSet)
	return res
}

// 定义：输入一棵二叉树，删除 delSet 中的节点，返回删除完成后的根节点
func doDelete(root *TreeNode, hasParent bool, res *[]*TreeNode, delSet map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}
	// 判断是否需要被删除
	deleted := delSet[root.Val]
	if !deleted && !hasParent {
		// 没有父节点且不需要被删除，就是一个新的根节点
		*res = append(*res, root)
	}
	// 去左右子树进行删除
	root.Left = doDelete(root.Left, !deleted, res, delSet)
	root.Right = doDelete(root.Right, !deleted, res, delSet)
	// 如果需要被删除，返回 nil 给父节点
	if deleted {
		return nil
	}
	return root
}

//原来的解法，理解错误，这个解法每次只删除一个节点，得到两棵树，需要把数组中的全部节点删除

//func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
//	res := make([]*TreeNode, 0)
//	for _, i := range to_delete {
//		res = append(res, helper1(root, i)...)
//	}
//	return res
//}
//
//func helper1(root *TreeNode, target int) []*TreeNode {
//	res := make([]*TreeNode, 0)
//	if root == nil {
//		return res
//	}
//	if root.Val == target {
//		if root.Left != nil {
//			res = append(res, root.Left)
//		}
//		if root.Right != nil {
//			res = append(res, root.Right)
//		}
//		return res
//	}
//	leftRes := helper1(root.Left, target)
//	rightRes := helper1(root.Right, target)
//	if len(leftRes) > 0 {
//		res = append(res, leftRes...)
//	}
//	if len(rightRes) > 0 {
//		res = append(res, rightRes...)
//	}
//	return res
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeleteNodesAndReturnForest(t *testing.T) {

}
