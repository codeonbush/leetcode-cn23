package leetcode

import (
	"testing"
)

/**
给你一棵二叉树的根节点 root ，返回树的 最大宽度 。

 树的 最大宽度 是所有层中最大的 宽度 。



 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节
点，这些 null 节点也计入长度。



 题目数据保证答案将会在 32 位 带符号整数范围内。



 示例 1：


输入：root = [1,3,2,5,3,null,9]
输出：4
解释：最大宽度出现在树的第 3 层，宽度为 4 (5,3,null,9) 。


 示例 2：


输入：root = [1,3,2,5,null,null,9,6,null,7]
输出：7
解释：最大宽度出现在树的第 4 层，宽度为 7 (6,null,null,null,null,null,7) 。


 示例 3：


输入：root = [1,3,2,5]
输出：2
解释：最大宽度出现在树的第 2 层，宽度为 2 (3,2) 。




 提示：


 树中节点的数目范围是 [1, 3000]
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 639 👎 0

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
//这道题的解题关键是要给二叉树节点按行进行编号，
//然后你就可以通过每一行的最左侧节点和最右侧节点的编号推算出这一行的宽度，进而算出最大宽度：
//而且，这样编号还可以通过父节点的编号计算得出左右子节点的编号：
//
//假设父节点的编号是 x，左子节点就是 2 * x，右子节点就是 2 * x + 1。
//
//这个特性常见于完全二叉树的题目当中，你可以去看看后文 图文详解二叉堆，实现优先级队列 或者去做一下 1104. 二叉树寻路 这道题。
//
//其中 DFS 的递归解法需要你透彻理解二叉树的遍历顺序，你可以先做一下 199. 二叉树的右视图 这道题。
//层序遍历思路
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    type Pair struct {
        node *TreeNode
        id   int
    }
    // 记录最大的宽度
    maxWidth := 0
    // 标准 BFS 层序遍历算法
    q := make([]Pair, 0)
    q = append(q, Pair{root, 1})
    // 从上到下遍历整棵树
    for len(q) > 0 {
        sz := len(q)
        start, end := 0, 0
        // 从左到右遍历每一行
        for i := 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            curNode, curId := cur.node, cur.id
            // 记录当前行第一个和最后一个节点的编号
            if i == 0 {
                start = curId
            }
            if i == sz-1 {
                end = curId
            }
            // 左右子节点入队，同时记录对应节点的编号
            if curNode.Left != nil {
                q = append(q, Pair{curNode.Left, curId * 2})
            }
            if curNode.Right != nil {
                q = append(q, Pair{curNode.Right, curId*2 + 1})
            }
        }
        // 用当前行的宽度更新最大宽度
        maxWidth = max(maxWidth, end - start + 1)
    }
    return maxWidth
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 递归遍历思路
//func widthOfBinaryTree2(root *TreeNode) int {
//    if root == nil {
//        return 0
//    }
//    maxWidth := 1
//    firstId := make([]int, 0)
//    traverse(root, 1, 1, &firstId, &maxWidth)
//    return maxWidth
//
//}
//
//// 二叉树遍历函数
//func traverse(root *TreeNode, id, depth int, firstId *[]int, maxWidth *int) {
//    if root == nil {
//        return
//    }
//
//    if len(*firstId) == depth-1 {
//        // 因为代码是先 traverse(root.left) 后 traverse(root.right)，
//        // 所以第一次到达这个深度一定是最左侧的节点，记录其编号
//        *firstId = append(*firstId, id)
//    } else {
//        // 这个深度的其他节点，负责计算更新当前深度的最大宽度
//        *maxWidth = max(*maxWidth, id-(*firstId)[depth-1]+1)
//    }
//
//    traverse(root.Left, id*2, depth+1, firstId, maxWidth)
//    traverse(root.Right, id*2+1, depth+1, firstId, maxWidth)
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumWidthOfBinaryTree(t *testing.T) {

}
