package leetcode

import (
	"testing"
)

/**
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。



 示例 1：




输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15


 示例 2：


输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：19




 提示：


 树中节点数目在范围 [1, 10⁴] 之间。
 1 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 174 👎 0

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
//key：按层遍历，最终得到的就是最深的一层
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	sum := 0
	// 使用 BFS 遍历树的每一层
	for len(q) != 0 {
		sum = 0
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			sum += cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	// 累加每一层的节点之和，最后得到的最深层节点的值和
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeepestLeavesSum(t *testing.T) {

}
