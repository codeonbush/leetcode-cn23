package leetcode

import (
	"testing"
)

/**
给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.



 示例 1：




输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]


 示例 2：




输入：root1 = [1,null,8], root2 = [8,1]
输出：[1,1,8,8]




 提示：


 每棵树的节点数在 [0, 5000] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 排序 👍 178 👎 0

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
// BST 有序迭代器
type BSTIterator struct {
	stk []*TreeNode
}

// 左侧树枝一撸到底
func (iter *BSTIterator) pushLeftBranch(p *TreeNode) {
	for p != nil {
		iter.stk = append(iter.stk, p)
		p = p.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.pushLeftBranch(root)
	return iter
}

func (iter *BSTIterator) Peek() int {
	return iter.stk[len(iter.stk)-1].Val
}

func (iter *BSTIterator) Next() int {
	p := iter.stk[len(iter.stk)-1]
	iter.stk = iter.stk[:len(iter.stk)-1]
	iter.pushLeftBranch(p.Right)
	return p.Val
}

func (iter *BSTIterator) HasNext() bool {
	return len(iter.stk) > 0
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// BST 有序迭代器
	t1 := Constructor(root1)
	t2 := Constructor(root2)
	res := []int{}
	// 类似合并有序链表的算法逻辑
	for t1.HasNext() && t2.HasNext() {
		if t1.Peek() > t2.Peek() {
			res = append(res, t2.Next())
		} else {
			res = append(res, t1.Next())
		}
	}
	// 如果有一棵 BST 还剩元素，添加到最后
	for t1.HasNext() {
		res = append(res, t1.Next())
	}
	for t2.HasNext() {
		res = append(res, t2.Next())
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAllElementsInTwoBinarySearchTrees(t *testing.T) {

}
