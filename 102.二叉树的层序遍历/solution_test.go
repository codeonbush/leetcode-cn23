package leetcode

import (
	"fmt"
	"testing"
)

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1715 👎 0
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelResult := make([]int, 0)
		for _, node := range queue {
			levelResult = append(levelResult, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelResult)
		queue = queue[len(levelResult):]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := levelOrder(root)
	fmt.Println(res)

	//testQueue := []int{1}
	//for _, i := range testQueue {
	//	fmt.Println(i)
	//	testQueue = append(testQueue, i+1)
	//}
	//fmt.Println("over")

	testQueue := []int{1}
	for i := 0; i < len(testQueue); i++ {
		fmt.Println(testQueue[i])
		testQueue = append(testQueue, testQueue[i]+1)
	}

}
