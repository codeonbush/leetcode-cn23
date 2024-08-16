package leetcode

import (
	"testing"
)

/**
完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。

 设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。

 实现 CBTInserter 类:


 CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
 CBTInserter.insert(int v) 向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态，并
返回插入节点 TreeNode 的父节点的值；
 CBTInserter.get_root() 将返回树的头节点。







 示例 1：




输入
["CBTInserter", "insert", "insert", "get_root"]
[[[1, 2]], [3], [4], []]
输出
[null, 1, 2, [1, 2, 3, 4]]

解释
CBTInserter cBTInserter = new CBTInserter([1, 2]);
cBTInserter.insert(3);  // 返回 1
cBTInserter.insert(4);  // 返回 2
cBTInserter.get_root(); // 返回 [1, 2, 3, 4]



 提示：


 树中节点数量范围为 [1, 1000]
 0 <= Node.val <= 5000
 root 是完全二叉树
 0 <= val <= 5000
 每个测试用例最多调用 insert 和 get_root 操作 10⁴ 次


 Related Topics 树 广度优先搜索 设计 二叉树 👍 168 👎 0

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

type CBTInserter struct {
	// 这个队列只记录完全二叉树底部可以进行插入的节点
	q    []*TreeNode
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	// 进行普通的 BFS，目的是找到底部可插入的节点
	temp := make([]*TreeNode, 0)
	temp = append(temp, root)

	q := make([]*TreeNode, 0)

	for len(temp) > 0 {
		cur := temp[0]
		temp = temp[1:]

		if cur.Left != nil {
			temp = append(temp, cur.Left)
		}

		if cur.Right != nil {
			temp = append(temp, cur.Right)
		}

		if cur.Right == nil || cur.Left == nil {
			// 找到完全二叉树底部可以进行插入的节点
			q = append(q, cur)
		}
	}

	return CBTInserter{q, root}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	cur := this.q[0]
	// 进行插入
	if cur.Left == nil {
		cur.Left = node
	} else if cur.Right == nil {
		cur.Right = node
		this.q = this.q[1:]
	}
	// 新节点的左右节点也是可以插入的
	this.q = append(this.q, node)
	return cur.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestCompleteBinaryTreeInserter(t *testing.T) {

}
