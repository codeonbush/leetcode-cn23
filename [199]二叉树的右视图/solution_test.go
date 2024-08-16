package leetcode

import (
	"testing"
)

/**
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。



 示例 1:




输入: [1,2,3,null,5,null,4]
输出: [1,3,4]


 示例 2:


输入: [1,null,3]
输出: [1,3]


 示例 3:


输入: []
输出: []




 提示:


 二叉树的节点个数的范围是 [0,100]

 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1056 👎 0

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

// 这题有两个思路：
//
// 1、用 BFS 层序遍历算法，每一层的最后一个节点就是二叉树的右侧视图。我们可以把 BFS 反过来，从右往左遍历每一行，进一步提升效率。
// 2、用 DFS 递归遍历算法，同样需要反过来，先递归 root.right 再递归 root.left，同时用 res 记录每一层的最右侧节点作为右侧视图。
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// BFS 层序遍历，计算右侧视图
	q := []*TreeNode{root}
	// while 循环控制从上向下一层层遍历
	for len(q) > 0 {
		//记录队列初始长度
		sz := len(q)
		// 每一层头部就是最右侧的元素
		last := q[0]
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			// 控制每一层从右向左遍历
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
		}
		// 每一层的最后一个节点就是二叉树的右侧视图
		res = append(res, last.Val)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeRightSideView(t *testing.T) {

}
