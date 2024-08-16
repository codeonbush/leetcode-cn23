package leetcode

import (
	"math"
	"sort"
	"testing"
)

/**
给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。

 对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结
点位于 (0, 0) 。

 二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结
点的值从小到大进行排序。

 返回二叉树的 垂序遍历 序列。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[9],[3,15],[20],[7]]
解释：
列 -1 ：只有结点 9 在此列中。
列  0 ：只有结点 3 和 15 在此列中，按从上到下顺序。
列  1 ：只有结点 20 在此列中。
列  2 ：只有结点 7 在此列中。

 示例 2：


输入：root = [1,2,3,4,5,6,7]
输出：[[4],[2],[1,5,6],[3],[7]]
解释：
列 -2 ：只有结点 4 在此列中。
列 -1 ：只有结点 2 在此列中。
列  0 ：结点 1 、5 和 6 都在此列中。
          1 在上面，所以它出现在前面。
          5 和 6 位置都是 (2, 0) ，所以按值从小到大排序，5 在 6 的前面。
列  1 ：只有结点 3 在此列中。
列  2 ：只有结点 7 在此列中。


 示例 3：


输入：root = [1,2,3,4,6,5,7]
输出：[[4],[2],[1,5,6],[3],[7]]
解释：
这个示例实际上与示例 2 完全相同，只是结点 5 和 6 在树中的位置发生了交换。
因为 5 和 6 的位置仍然相同，所以答案保持不变，仍然按值从小到大排序。



 提示：


 树中结点数目总数在范围 [1, 1000] 内
 0 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 排序 👍 295 👎 0

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

// 记录每个节点和对应的坐标 (row, col)
type Triple struct {
	row  int
	col  int
	node *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	var nodes []*Triple

	// 二叉树遍历函数，记录所有节点对应的坐标
	var traverse func(*TreeNode, int, int)
	traverse = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}
		// 记录坐标
		nodes = append(nodes, &Triple{row: row, col: col, node: root})
		// 二叉树遍历框架
		traverse(root.Left, row+1, col-1)
		traverse(root.Right, row+1, col+1)
	}

	// 遍历二叉树，并且为所有节点生成对应的坐标
	traverse(root, 0, 0)

	// 根据题意，根据坐标值对所有节点进行排序：
	// 按照 col 从小到大排序，col 相同的话按 row 从小到大排序，
	// 如果 col 和 row 都相同，按照 node.val 从小到大排序。
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col == nodes[j].col && nodes[i].row == nodes[j].row {
			return nodes[i].node.Val < nodes[j].node.Val
		}
		if nodes[i].col == nodes[j].col {
			return nodes[i].row < nodes[j].row
		}
		return nodes[i].col < nodes[j].col
	})

	// 将排好序的节点组装成题目要求的返回格式
	res := [][]int{}
	// 记录上一列编号，初始化一个特殊值
	preCol := math.MinInt32
	for i := 0; i < len(nodes); i++ {
		cur := nodes[i]
		if cur.col != preCol {
			// 开始记录新的一列
			res = append(res, []int{})
			preCol = cur.col
		}
		res[len(res)-1] = append(res[len(res)-1], cur.node.Val)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestVerticalOrderTraversalOfABinaryTree(t *testing.T) {

}
