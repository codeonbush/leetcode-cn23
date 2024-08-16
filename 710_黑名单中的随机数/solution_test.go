package leetcode

import (
	"math/rand"
	"testing"
	"time"
)

/**
给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入 黑名单
blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。

 优化你的算法，使它最小化调用语言 内置 随机函数的次数。

 实现 Solution 类:


 Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
 int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数




 示例 1：


输入
["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
[[7, [2, 3, 5]], [], [], [], [], [], [], []]
输出
[null, 0, 4, 1, 6, 1, 0, 4]

解释
Solution solution = new Solution(7, [2, 3, 5]);
solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
                 // 0、1、4和6的返回概率必须相等(即概率为1/4)。
solution.pick(); // 返回 4
solution.pick(); // 返回 1
solution.pick(); // 返回 6
solution.pick(); // 返回 1
solution.pick(); // 返回 0
solution.pick(); // 返回 4




 提示:


 1 <= n <= 10⁹
 0 <= blacklist.length <= min(10⁵, n - 1)
 0 <= blacklist[i] < n
 blacklist 中所有值都 不同
 pick 最多被调用 2 * 10⁴ 次


 Related Topics 数组 哈希表 数学 二分查找 排序 随机化 👍 241 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// Solution struct
type Solution struct {
	sz      int
	mapping map[int]int
	rand    *rand.Rand
}

// Constructor function
func Constructor(N int, blacklist []int) Solution {
	sz := N - len(blacklist)
	mapping := make(map[int]int)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 转化成集合，方便快速查询
	blackSet := make(map[int]bool)
	for _, b := range blacklist {
		blackSet[b] = true
	}

	last := N - 1
	for _, b := range blacklist {
		// 如果 b 已经在区间 [sz, N)，可以直接忽略
		if b >= sz {
			continue
		}
		// 跳过所有黑名单中的数字
		for blackSet[last] {
			last--
		}
		// 将黑名单中的索引映射到合法数字
		mapping[b] = last
		last--
	}

	return Solution{
		sz:      sz,
		mapping: mapping,
		rand:    rand,
	}
}

// pick function
func (this *Solution) Pick() int {
	// 随机选取一个索引
	index := this.rand.Intn(this.sz)
	// 这个索引命中了黑名单，需要被映射到其他位置
	if val, found := this.mapping[index]; found {
		return val
	}
	// 若没命中黑名单，则直接返回
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRandomPickWithBlacklist(t *testing.T) {

}
