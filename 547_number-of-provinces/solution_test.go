package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**


 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。




 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
isConnected[i][j] = 0 表示二者不直接相连。

 返回矩阵中 省份 的数量。



 示例 1：


输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2


 示例 2：


输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3




 提示：


 1 <= n <= 200
 n == isConnected.length
 n == isConnected[i].length
 isConnected[i][j] 为 1 或 0
 isConnected[i][i] == 1
 isConnected[i][j] == isConnected[j][i]


 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 1172 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUF(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 如果两个城市是相连的，那么它们属于同一个连通分量
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.getCount()
}

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
func (uf *UF) union(p, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)

	if rootP == rootQ {
		return
	}

	uf.parent[rootQ] = rootP
	// 两个连通分量合并成一个连通分量
	uf.count--
}

// 判断节点 p 和节点 q 是否连通
func (uf *UF) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// 返回图中的连通分量个数
func (uf *UF) getCount() int {
	return uf.count
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfProvinces(t *testing.T) {
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
