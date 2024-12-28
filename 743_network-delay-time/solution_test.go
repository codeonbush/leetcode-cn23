package leetcode

import (
	"container/heap"
	"fmt"
	"math"
	"reflect"
	"testing"
)

/**
有 n 个网络节点，标记为 1 到 n。

 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi
是一个信号从源节点传递到目标节点的时间。

 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。



 示例 1：




输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2


 示例 2：


输入：times = [[1,2,1]], n = 2, k = 1
输出：1


 示例 3：


输入：times = [[1,2,1]], n = 2, k = 2
输出：-1




 提示：


 1 <= k <= n <= 100
 1 <= times.length <= 6000
 times[i].length == 3
 1 <= ui, vi <= n
 ui != vi
 0 <= wi <= 100
 所有 (ui, vi) 对都 互不相同（即，不含重复边）


 Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 830 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func networkDelayTime(times [][]int, n int, k int) int {
	// 构建图的邻接表
	graph := make([][]*pair, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]*pair, 0)
	}
	// 填充图的邻接表
	for _, edge := range times {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], &pair{weight, to})
	}

	// 使用Dijkstra算法计算从节点k到其他节点的最短路径
	distTo := dijkstra(k, graph, n)

	// 找到最长的最短路径
	res := 0
	for i := 1; i <= n; i++ {
		if distTo[i] == math.MaxInt32 {
			return -1 // 如果有节点不可达，返回-1
		}
		if distTo[i] > res {
			res = distTo[i]
		}
	}
	return res
}

// Dijkstra算法实现
func dijkstra(start int, graph [][]*pair, n int) []int {
	distTo := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pair{0, start})

	for pq.Len() > 0 {
		top := heap.Pop(pq).(*pair)
		curNode := top.index
		curDist := top.value

		if curDist > distTo[curNode] {
			continue
		}

		for _, neighbor := range graph[curNode] {
			nextNode := neighbor.index
			nextDist := curDist + neighbor.value

			if nextDist < distTo[nextNode] {
				distTo[nextNode] = nextDist
				heap.Push(pq, &pair{nextDist, nextNode})
			}
		}
	}
	return distTo
}

// pair结构体用于存储节点和距离
type pair struct {
	value int
	index int
}

// 优先队列的实现
type priorityQueue []*pair

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*pair))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNetworkDelayTime(t *testing.T) {
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
