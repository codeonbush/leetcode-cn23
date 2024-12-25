package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。

 如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。

 给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。



 示例 1：


输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
输出：5
解释：一种移除 5 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
石头 [0,0] 不能移除，因为它没有与另一块石头同行/列。

 示例 2：


输入：stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
输出：3
解释：一种移除 3 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,0] 同行。
2. 移除石头 [2,0] ，因为它和 [0,0] 同列。
3. 移除石头 [0,2] ，因为它和 [0,0] 同行。
石头 [0,0] 和 [1,1] 不能移除，因为它们没有与另一块石头同行/列。

 示例 3：


输入：stones = [[0,0]]
输出：0
解释：[0,0] 是平面上唯一一块石头，所以不可以移除它。



 提示：


 1 <= stones.length <= 1000
 0 <= xi, yi <= 10⁴
 不会有两块石头放在同一个坐标点上


 Related Topics 深度优先搜索 并查集 图 哈希表 👍 369 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func removeStones(stones [][]int) int {
	n := len(stones)

	// 一维坐标 -> 节点 ID
	codeToId := make(map[int]int)
	for i := 0; i < n; i++ {
		codeToId[encode(stones[i])] = i
	}

	// 记录每一行每一列有哪些节点
	colIndexToCodes := make(map[int][]int)
	rowIndexToCodes := make(map[int][]int)
	for _, point := range stones {
		x, y := point[0], point[1]
		if _, ok := rowIndexToCodes[x]; !ok {
			rowIndexToCodes[x] = []int{}
		}
		if _, ok := colIndexToCodes[y]; !ok {
			colIndexToCodes[y] = []int{}
		}
		rowIndexToCodes[x] = append(rowIndexToCodes[x], encode(point))
		colIndexToCodes[y] = append(colIndexToCodes[y], encode(point))
	}

	// 启动 union find 算法
	uf := NewUF(n)

	// 把每一列的节点连通
	for _, col := range colIndexToCodes {
		firstId := codeToId[col[0]]
		for i := 1; i < len(col); i++ {
			otherId := codeToId[col[i]]
			uf.Union(firstId, otherId)
		}
	}

	// 把每一行的节点连通
	for _, row := range rowIndexToCodes {
		firstId := codeToId[row[0]]
		for i := 1; i < len(row); i++ {
			otherId := codeToId[row[i]]
			uf.Union(firstId, otherId)
		}
	}
	// 石头总数 - 连通分量数量就是被消除的石头个数
	return n - uf.Count()
}

// 将二维坐标转化成一维索引
func encode(point []int) int {
	return point[0]*10000 + point[1]
}

// union find 算法模板
type UF struct {
	// 连通分量个数
	count int
	// 存储每个节点的父节点
	parent []int
}

// n 为图中节点的个数
func NewUF(n int) *UF {
	uf := &UF{
		count:  n,
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

// 将节点 p 和节点 q 连通
func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	uf.parent[rootQ] = rootP
	// 两个连通分量合并成一个连通分量
	uf.count--
}

// 判断节点 p 和节点 q 是否连通
func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// 返回图中的连通分量个数
func (uf *UF) Count() int {
	return uf.count
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMostStonesRemovedWithSameRowOrColumn(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{10, 6, 8, 5, 11, 9},
			expected: []int{3, 1, 1, 1, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := longestWPI(test.input)
			fmt.Println(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For heights %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
