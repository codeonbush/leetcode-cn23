package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。 

 你可以按 任何顺序 返回答案。 

 

 示例 1： 

 
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
] 

 示例 2： 

 
输入：n = 1, k = 1
输出：[[1]] 

 

 提示： 

 
 1 <= n <= 20 
 1 <= k <= n 
 

 Related Topics 回溯 👍 1711 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var res [][]int
	// 记录回溯算法的递归路径
	var track []int

	// 回溯算法核心函数
	var backtrack func(int)
	backtrack = func(start int) {
		// base case
		if k == len(track) {
			// 遍历到了第 k 层，收集当前节点的值
			res = append(res, append([]int(nil), track...))
			return
		}

		// 回溯算法标准框架
		for i := start; i <= n; i++ {
			// 选择
			track = append(track, i)
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集
			backtrack(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func TestCombinations(t *testing.T){
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