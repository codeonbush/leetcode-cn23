package leetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/**
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

 连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示
val 的绝对值。

 请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。



 示例 1：




输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
输出：20
解释：

我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
注意到任意两个点之间只有唯一一条路径互相到达。


 示例 2：


输入：points = [[3,12],[-2,5],[-4,1]]
输出：18


 示例 3：


输入：points = [[0,0],[1,1],[1,0],[-1,1]]
输出：4


 示例 4：


输入：points = [[-1000000,-1000000],[1000000,1000000]]
输出：4000000


 示例 5：


输入：points = [[0,0]]
输出：0




 提示：


 1 <= points.length <= 1000
 -10⁶ <= xi, yi <= 10⁶
 所有点 (xi, yi) 两两不同。


 Related Topics 并查集 图 数组 最小生成树 👍 334 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minCostConnectPoints(points [][]int) int {
	n := len(points)

	// 生成所有边及权重
	edges := make([][]int, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			// 用坐标点在 points 中的索引表示坐标点
			edges = append(edges, []int{
				i, j, abs(xi-xj) + abs(yi-yj),
			})
		}
	}

	// 将边按照权重从小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// 执行 Kruskal 算法
	mst := 0
	uf := NewUF(n)
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		// 若这条边会产生环，则不能加入 mst
		if uf.Connected(u, v) {
			continue
		}
		// 若这条边不会产生环，则属于最小生成树
		mst += weight
		uf.Union(u, v)
	}
	return mst
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func TestMinCostToConnectAllPoints(t *testing.T) {
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
