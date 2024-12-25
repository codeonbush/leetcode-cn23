package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=
b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。 

 只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。 

 

 
 

 示例 1： 

 输入：["a==b","b!=a"]
输出：false
解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
 

 示例 2： 

 输入：["b==a","a==b"]
输出：true
解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
 

 示例 3： 

 输入：["a==b","b==c","a==c"]
输出：true
 

 示例 4： 

 输入：["a==b","b!=c","c==a"]
输出：false
 

 示例 5： 

 输入：["c==c","b==d","x!=z"]
输出：true
 

 

 提示： 

 
 1 <= equations.length <= 500 
 equations[i].length == 4 
 equations[i][0] 和 equations[i][3] 是小写字母 
 equations[i][1] 要么是 '='，要么是 '!' 
 equations[i][2] 是 '=' 
 

 Related Topics 并查集 图 数组 字符串 👍 336 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func equationsPossible(equations []string) bool {
	// 26 个英文字母
	uf := NewUF(26)
	// 先让相等的字母形成连通分量
	for _, eq := range equations {
		if eq[1] == '=' {
			x := eq[0]
			y := eq[3]
			uf.Union(int(x - 'a'), int(y - 'a'))
		}
	}
	// 检查不等关系是否打破相等关系的连通性
	for _, eq := range equations {
		if eq[1] == '!' {
			x := eq[0]
			y := eq[3]
			// 如果相等关系成立，就是逻辑冲突
			if uf.Connected(int(x - 'a'), int(y - 'a')) {
				return false
			}
		}
	}
	return true
}

type UF struct {
	// 连通分量个数
	count int
	// 存储每个节点的父节点
	parent []int
}

// n 为图中节点的个数
func NewUF(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{
		count:  n,
		parent: parent,
	}
}

// 将节点 p 和节点 q 连通
func (u *UF) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	u.parent[rootQ] = rootP
	// 两个连通分量合并成一个连通分量
	u.count--
}

// 判断节点 p 和节点 q 是否连通
func (u *UF) Connected(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	return rootP == rootQ
}

func (u *UF) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

// 返回图中的连通分量个数
func (u *UF) Count() int {
	return u.count
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSatisfiabilityOfEqualityEquations(t *testing.T){
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