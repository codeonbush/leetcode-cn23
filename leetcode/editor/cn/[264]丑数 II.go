package leetcode

import (
	"testing"
)

/**
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

 丑数 就是质因子只包含 2、3 和 5 的正整数。



 示例 1：


输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。


 示例 2：


输入：n = 1
输出：1
解释：1 通常被视为丑数。




 提示：


 1 <= n <= 1690


 Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 1190 👎 0

*/
//类似 如何高效寻找素数 的思路：如果一个数 x 是丑数，那么 x * 2, x * 3, x * 5 都一定是丑数。
//
//我们把所有丑数想象成一个从小到大排序的链表，就是这个样子：
//
//1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 8 -> ...
//然后，我们可以把丑数分为三类：2 的倍数、3 的倍数、5 的倍数（按照题目的意思，1 算作特殊的丑数，放在开头），这三类丑数就好像三条有序链表，如下：
//
//能被 2 整除的丑数：
//
//1 -> 1*2 -> 2*2 -> 3*2 -> 4*2 -> 5*2 -> 6*2 -> 8*2 ->...
//能被 3 整除的丑数：
//
//1 -> 1*3 -> 2*3 -> 3*3 -> 4*3 -> 5*3 -> 6*3 -> 8*3 ->...
//能被 5 整除的丑数：
//
//1 -> 1*5 -> 2*5 -> 3*5 -> 4*5 -> 5*5 -> 6*5 -> 8*5 ->...
//我们其实就是想把这三条「有序链表」合并在一起并去重，合并的结果就是丑数的序列。
//然后求合并后的这条有序链表中第 n 个元素是什么。所以这里就和 链表双指针技巧汇总 中讲到的合并 k 条有序链表的思路基本一样了。
//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	// 可以理解为三个指向有序链表头结点的指针
	p2, p3, p5 := 1, 1, 1
	// 可以理解为三个有序链表的头节点的值
	product2, product3, product5 := 1, 1, 1
	// 可以理解为最终合并的有序链表（结果链表）
	ugly := make([]int, n+1)
	// 可以理解为结果链表上的指针
	p := 1

	// 开始合并三个有序链表
	for p <= n {
		// 取三个链表的最小结点
		min := min(min(product2, product3), product5)
		// 接到结果链表上
		ugly[p] = min
		p++
		// 前进对应有序链表上的指针
		if min == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if min == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if min == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}
	}
	// 返回第 n 个丑数
	return ugly[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUglyNumberIi(t *testing.T) {

}
