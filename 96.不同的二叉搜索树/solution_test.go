package leetcode

import (
	"testing"
)

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2277 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

// 组合数=左子树的组合数 x 右子树的组合数
func numTrees(n int) int {
	mem := make([][]int, n+1)
	for i, _ := range mem {
		mem[i] = make([]int, n+1)
	}
	var count func(l, r int) int
	count = func(l, r int) int {
		//当 lo > hi 闭区间 [lo, hi] 是个空区间，也就对应着空节点 null
		if l > r {
			return 1
		}
		if mem[l][r] != 0 {
			return mem[l][r]
		}
		result := 0
		for i := l; i <= r; i++ {
			left := count(l, i-1)
			right := count(i+1, r)
			result += left * right
		}
		mem[l][r] = result
		return result
	}
	return count(1, n)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUniqueBinarySearchTrees(t *testing.T) {

}
