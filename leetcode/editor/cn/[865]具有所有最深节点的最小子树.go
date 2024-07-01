package leetcode

import (
	"testing"
)

/**
给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。

 返回包含原始树中所有 最深节点 的 最小子树 。

 如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。

 一个节点的 子树 是该节点加上它的所有后代的集合。



 示例 1：




输入：root = [3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：
我们返回值为 2 的节点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的节点。
注意，节点 5、3 和 2 包含树中最深的节点，但节点 2 的子树最小，因此我们返回它。


 示例 2：


输入：root = [1]
输出：[1]
解释：根节点是树中最深的节点。

 示例 3：


输入：root = [0,1,3,null,2]
输出：[2]
解释：树中最深的节点为 2 ，有效子树为节点 2、1 和 0 的子树，但节点 2 的子树最小。



 提示：


 树中节点的数量在
 [1, 500] 范围内。
 0 <= Node.val <= 500
 每个节点的值都是 独一无二 的。




 注意：本题与力扣 1123 重复：https://leetcode-cn.com/problems/lowest-common-ancestor-of-
deepest-leaves

 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 220 👎 0

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
//二叉树的递归分为「遍历」和「分解问题」两种思维模式，这道题需要用到「分解问题」的思维，而且涉及处理子树，需要用后序遍历。
//
//说到底，这道题就是让你求那些「最深」的叶子节点的最近公共祖先，可以看下前文 二叉树最近公共祖先。
//
//你想想，一个节点需要知道哪些信息，才能确定自己是最深叶子节点的最近公共祖先？
//
//它需要知道自己的左右子树的最大深度：如果左右子树一样深，那么当前节点就是最近公共祖先；如果左右子树不一样深，那么最深叶子节点的最近公共祖先肯定在左右子树上。
//
//所以我们新建一个 Result 类，存放左右子树的最大深度及叶子节点的最近公共祖先节点，其余逻辑类似 104. 二叉树的最大深度。
type Result struct {
	node  *TreeNode
	depth int
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res := maxDepth(root)
	return res.node
}

// 定义：输入一棵二叉树，返回该二叉树的最大深度以及最深叶子节点的最近公共祖先节点
func maxDepth(root *TreeNode) Result {
	if root == nil {
		return Result{}
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left.depth == right.depth {
		// 当左右子树的最大深度相同时，这个根节点是新的最近公共祖先
		// 以当前 root 节点为根的子树深度是子树深度 + 1
		return Result{root, left.depth + 1}
	}
	// 左右子树的深度不同，则最近公共祖先在 depth 较大的一边
	res := Result{}
	if left.depth > right.depth {
		res = left
	} else {
		res = right
	}
	// 正确维护二叉树的最大深度
	res.depth++

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSmallestSubtreeWithAllTheDeepestNodes(t *testing.T) {

}
