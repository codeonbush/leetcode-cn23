package leetcode

import(
    "testing"
)

//给你一棵二叉树的根节点，返回该树的 直径 。 
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。 
//
// 两节点之间路径的 长度 由它们之间边数表示。 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [1,2,3,4,5]
//输出：3
//解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 
//
// 示例 2： 
//
// 
//输入：root = [1,2]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [1, 10⁴] 内 
// -100 <= Node.val <= 100 
// 
//
// Related Topics 树 深度优先搜索 二叉树 👍 1446 👎 0


Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
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

//直径=左子树最大深度+右子树最大深度
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    var maxDepth func(root *TreeNode) int
    maxDepth = func(root *TreeNode) int {
        if root == nil{
            return 0;
        }
        maxDepthL := maxDepth(root.Left)
        maxDepthR := maxDepth(root.Right)
        maxDiameter = max(maxDiameter, maxDepthL + maxDepthR)
        return max(maxDepthL, maxDepthR)+ 1;
    }
    maxDepth(root)
    return maxDiameter
}
//leetcode submit region end(Prohibit modification and deletion)


func TestDiameterOfBinaryTree(t *testing.T){
    //cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d expected %d, but %d got",
	//				c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}
}