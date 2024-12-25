package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
二叉树上有 n 个节点，按从 0 到 n - 1 编号，其中节点 i 的两个子节点分别是 leftChild[i] 和 rightChild[i]。 

 只有 所有 节点能够形成且 只 形成 一颗 有效的二叉树时，返回 true；否则返回 false。 

 如果节点 i 没有左子节点，那么 leftChild[i] 就等于 -1。右子节点也符合该规则。 

 注意：节点没有值，本问题中仅仅使用节点编号。 

 

 示例 1： 

 

 
输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
输出：true
 

 示例 2： 

 

 
输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
输出：false
 

 示例 3： 

 

 
输入：n = 2, leftChild = [1,0], rightChild = [-1,-1]
输出：false
 

 

 提示： 

 
 n == leftChild.length == rightChild.length 
 1 <= n <= 10⁴ 
 -1 <= leftChild[i], rightChild[i] <= n - 1 
 

 Related Topics 树 深度优先搜索 广度优先搜索 并查集 图 二叉树 👍 133 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
// 用 Union Find 算法判断
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// 记录每个节点的入度
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			indegree[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			indegree[rightChild[i]]++
		}
	}
	// 按道理应该有且只有根节点的入度为 0，
	// 其他节点的入度都必须为 1
	root := -1
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			if root != -1 {
				// 有多个入度为 0 的节点
				return false
			}
			root = i
		} else if indegree[i] != 1 {
			// 除了根节点外其他节点的入度都必须为 1
			return false
		}
	}

	// 如果没有根节点，那肯定不是合法二叉树
	if root == -1 {
		return false
	}

	// 启动 Union-Find 并查集算法，
	// 保证树中只有一个联通分量且不成环
	uf := newUF(n)
	for i := 0; i < n; i++ {
		left := leftChild[i]
		right := rightChild[i]

		if left != -1 {
			if uf.connected(i, left) {
				// 成环
				return false
			}
			uf.union(i, left)
		}
		if right != -1 {
			if uf.connected(i, right) {
				// 成环
				return false
			}
			uf.union(i, right)
		}
	}
	// 要保证只有一个连通分量
	return uf.count == 1
}

type UF struct {
	// 记录连通分量个数
	count int
	// 存储若干棵树
	parent []int
	// 记录树的“重量”
	size []int
}

func newUF(n int) *UF {
	uf := &UF{
		count:  n,
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// 将 p 和 q 连通
func (uf *UF) union(p, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}

	// 小树接到大树下面，较平衡
	if uf.size[rootP] > uf.size[rootQ] {
		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	} else {
		uf.parent[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	}
	uf.count--
}

// 判断 p 和 q 是否互相连通
func (uf *UF) connected(p, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	// 处于同一棵树上的节点，相互连通
	return rootP == rootQ
}

// 返回节点 x 的根节点
func (uf *UF) find(x int) int {
	for uf.parent[x] != x {
		// 进行路径压缩
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)


func TestValidateBinaryTreeNodes(t *testing.T){
    tests := []struct {
		input  []int
		expected []int
	}{
		{
			input:  []int{10, 6, 8, 5, 11, 9},
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