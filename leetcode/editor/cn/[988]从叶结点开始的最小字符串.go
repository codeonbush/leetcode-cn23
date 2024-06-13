package leetcode

import (
	"math"
	"testing"
)

/**
给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'。

 返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。


 注：字符串中任何较短的前缀在 字典序上 都是 较小 的：



 例如，在字典序上 "ab" 比 "aba" 要小。叶结点是指没有子结点的结点。


 节点的叶节点是没有子节点的节点。






 示例 1：




输入：root = [0,1,2,3,4,3,4]
输出："dba"


 示例 2：




输入：root = [25,1,3,1,3,0,2]
输出："adz"


 示例 3：




输入：root = [2,2,1,null,1,0,null,0]
输出："abc"




 提示：


 给定树的结点数在 [1, 8500] 范围内
 0 <= Node.val <= 25


 Related Topics 树 深度优先搜索 字符串 二叉树 👍 109 👎 0

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
//func smallestFromLeaf(root *TreeNode) string {
//	// 遍历过程中的路径
//	path := new(strings.Builder)
//	var dfs func(*TreeNode)
//	res := ""
//
//	dfs = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		if root.Left == nil && root.Right == nil {
//			// 找到叶子结点，比较字典序最小的路径
//			// 结果字符串是从叶子向根，所以需要反转
//			path.WriteRune(rune('a' + root.Val))
//			defer path.Reset()
//
//			s := path.String()
//			if res == "" || res > s {
//				// 如果字典序更小，则更新 res
//				res = s
//			}
//			return
//		}
//		// 前序位置
//		path.WriteRune(rune('a' + root.Val))
//
//		dfs(root.Left)
//		dfs(root.Right)
//
//		// 后序位置
//		defer path.Reset()
//	}
//
//	dfs(root)
//	return res
//}

func smallestFromLeaf(root *TreeNode) string {
	// 遍历过程中的路径
	path := make([]rune, 0)
	var dfs func(*TreeNode)
	res := string(rune(math.MaxInt32)) // 使用一个大的初始值

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 将当前节点加入路径
		path = append(path, rune('a'+node.Val))

		if node.Left == nil && node.Right == nil {
			// 找到叶子结点，比较字典序最小的路径
			// 结果字符串是从根向叶子，所以需要在这里反转当前路径
			s := reverse(path)
			if res > s {
				// 如果字典序更小，则更新 res
				res = s
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}

		// 移除当前节点，回溯到父节点
		path = path[:len(path)-1]
	}

	dfs(root)
	return res
}

// 反转 rune 切片并返回字符串
func reverse(runes []rune) string {
	n := len(runes)
	reversed := make([]rune, n)
	for i, r := range runes {
		reversed[n-1-i] = r
	}
	return string(reversed)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSmallestStringStartingFromLeaf(t *testing.T) {

}
