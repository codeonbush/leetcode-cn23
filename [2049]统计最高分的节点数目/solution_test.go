package leetcode

import (
	"testing"
)

/**
给你一棵根节点为 0 的 二叉树 ，它总共有 n 个节点，节点编号为 0 到 n - 1 。同时给你一个下标从 0 开始的整数数组 parents 表示这棵树，
其中 parents[i] 是节点 i 的父节点。由于节点 0 是根，所以 parents[0] == -1 。

 一个子树的 大小 为这个子树内节点的数目。每个节点都有一个与之关联的 分数 。求出某个节点分数的方法是，将这个节点和与它相连的边全部 删除 ，剩余部分是若干个
 非空 子树，这个节点的 分数 为所有这些子树 大小的乘积 。

 请你返回有 最高得分 节点的 数目 。



 示例 1:



 输入：parents = [-1,2,0,2,0]
输出：3
解释：
- 节点 0 的分数为：3 * 1 = 3
- 节点 1 的分数为：4 = 4
- 节点 2 的分数为：1 * 1 * 2 = 2
- 节点 3 的分数为：4 = 4
- 节点 4 的分数为：4 = 4
最高得分为 4 ，有三个节点得分为 4 （分别是节点 1，3 和 4 ）。


 示例 2：



 输入：parents = [-1,2,0]
输出：2
解释：
- 节点 0 的分数为：2 = 2
- 节点 1 的分数为：2 = 2
- 节点 2 的分数为：1 * 1 = 1
最高分数为 2 ，有两个节点分数为 2 （分别为节点 0 和 1 ）。




 提示：


 n == parents.length
 2 <= n <= 10⁵
 parents[0] == -1
 对于 i != 0 ，有 0 <= parents[i] <= n - 1
 parents 表示一棵二叉树。


 Related Topics 树 深度优先搜索 数组 二叉树 👍 161 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 在做这道题之前，建议你先去看下我给 1339. 分裂二叉树的最大乘积 写的思路和解法代码，然后立马就知道这道题的思路了。
//
// 简单说，一个节点的 分数 = 左子树节点个数 x 右子树节点个数 x 除自己外其他节点个数。
//
// 只要写个 countNode 函数，在后序位置可以得到左右子树的节点个数 leftCount 和 rightCount，
//
// 然后除自己外其他节点个数 otherCount 就等于总的节点个数 n 减掉左右子树的节点个数再减掉当前节点，最后求个乘积就能算出当前节点的「分数」了。
//
// 当然，这道题还有个难点就是：题目给的 parents 数组不是我们经常见到的二叉树形式。
//
// 但问题不大，我们可以把 parents 数组转化成类似 图论基础 中讲到的邻接表结构，然后就可以像操作常规二叉树一样写算法了。
func countHighestScoreNodes(parents []int) int {
	// 用邻接表表示的一棵二叉树
	tree := buildTree(parents)

	scoreToCount := make(map[int64]int)

	// 计算二叉树中的节点个数
	var countNode func(root int) int
	countNode = func(root int) int {
		if root == -1 {
			return 0
		}
		// 二叉树中节点总数
		n := len(tree)
		leftCount := countNode(tree[root][0])
		rightCount := countNode(tree[root][1])

		// 后序位置，计算每个节点的「分数」
		otherCount := n - leftCount - rightCount - 1
		// 注意，这里要把 int 转化成 int64，否则会产生溢出！！！
		score := int64(max(leftCount, 1) *
			max(rightCount, 1) *
			max(otherCount, 1))
		// 给分数 score 计数
		scoreToCount[score]++

		return leftCount + rightCount + 1
	}

	countNode(0)

	// 计算最大分数出现的次数
	var maxScore int64 = 0
	for score := range scoreToCount {
		if score > maxScore {
			maxScore = score
		}
	}
	return scoreToCount[maxScore]
}

// 将 parents 数组转化成常规二叉树（邻接表形式）
func buildTree(parents []int) [][]int {
	n := len(parents)
	// 表节点 x 的左子节点为 tree[x][0]，节点 x 的右子节点为 tree[x][1]
	// 若 tree[x][0] 或 tree[x][1] 等于 -1 则代表空指针
	tree := make([][]int, n)
	for i := 0; i < n; i++ {
		// 先都初始化成空指针
		tree[i] = []int{-1, -1}
	}
	// 根据 parents 数组构建二叉树（跳过 parents[0] 根节点）
	for i := 1; i < n; i++ {
		parent_i := parents[i]
		if tree[parent_i][0] == -1 {
			tree[parent_i][0] = i
		} else {
			tree[parent_i][1] = i
		}
	}
	return tree
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountNodesWithTheHighestScore(t *testing.T) {

}
