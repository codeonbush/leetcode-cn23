package leetcode

import (
	"testing"
)

/**
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k ，返回到目标结点 target 距离为 k 的所有结点的值的数组。

 答案可以以 任何顺序 返回。






 示例 1：




输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
输出：[7,4,1]
解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1


 示例 2:


输入: root = [1], target = 1, k = 3
输出: []




 提示:


 节点数在 [1, 500] 范围内
 0 <= Node.val <= 500
 Node.val 中所有值 不同
 目标结点 target 是树上的结点。
 0 <= k <= 1000




 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 697 👎 0

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
//这道题常规的解法就是把二叉树变成一幅「图」，
//然后在图中用 BFS 算法求距离 target 节点 k 步的所有节点。
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 记录父节点：node.val -> parentNode
	// 题目说了树中所有节点值都是唯一的，所以可以用 node.val 代表 TreeNode
	parent := make(map[int]*TreeNode)

	var traverse func(*TreeNode, *TreeNode)
	traverse = func(root *TreeNode, parentNode *TreeNode) {
		if root == nil {
			return
		}
		parent[root.Val] = parentNode
		// 二叉树递归框架
		traverse(root.Left, root)
		traverse(root.Right, root)
	}

	traverse(root, nil)

	// 开始从 target 节点施放 BFS 算法，找到距离为 k 的节点
	q := make([]*TreeNode, 0)
	visited := make(map[int]bool)
	q = append(q, target)
	visited[target.Val] = true
	// 记录离 target 的距离
	dist := 0
	res := make([]int, 0)

	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if dist == k {
				// 找到距离起点 target 距离为 k 的节点
				res = append(res, cur.Val)
			}
			// 向父节点、左右子节点扩散
			parentNode := parent[cur.Val]
			if parentNode != nil && !visited[parentNode.Val] {
				visited[parentNode.Val] = true
				q = append(q, parentNode)
			}
			if cur.Left != nil && !visited[cur.Left.Val] {
				visited[cur.Left.Val] = true
				q = append(q, cur.Left)
			}
			if cur.Right != nil && !visited[cur.Right.Val] {
				visited[cur.Right.Val] = true
				q = append(q, cur.Right)
			}
		}
		// 向外扩展一圈
		dist++
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAllNodesDistanceKInBinaryTree(t *testing.T) {

}
