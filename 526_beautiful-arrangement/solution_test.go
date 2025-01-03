package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ： 

 
 perm[i] 能够被 i 整除 
 i 能够被 perm[i] 整除 
 

 给你一个整数 n ，返回可以构造的 优美排列 的 数量 。 

 

 示例 1： 

 
输入：n = 2
输出：2
解释：
第 1 个优美的排列是 [1,2]：
    - perm[1] = 1 能被 i = 1 整除
    - perm[2] = 2 能被 i = 2 整除
第 2 个优美的排列是 [2,1]:
    - perm[1] = 2 能被 i = 1 整除
    - i = 2 能被 perm[2] = 1 整除
 

 示例 2： 

 
输入：n = 1
输出：1
 

 

 提示： 

 
 1 <= n <= 15 
 

 Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 416 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译。
// 本代码的正确性已通过力扣验证，如有疑问，可以对照 java 代码查看。

func countArrangement(n int) int {
	// 记录所有的「优美排列」的个数
	used := make([]bool, n+1)
	var res int
	// 记录回溯算法的递归路径，即每个索引选择的元素
	var track []int
	// track 中的元素会被标记为 true，避免重复选择
	backtrack(n, 1, &used, &res, &track)
	return res
}

// 回溯算法标准框架，站在索引的视角选择元素
func backtrack(n int, index int, used *[]bool, res *int, track *[]int) {
	// base case，到达叶子节点
	if index > n {
		// 找到一个结果
		*res += 1
		return
	}

	// 索引 index 开始选择元素
	for elem := 1; elem <= n; elem++ {
		// 已经被其他索引选过的元素，不能重复选择
		if (*used)[elem] {
			continue
		}
		if !(index%elem == 0 || elem%index == 0) {
			continue
		}
		// 做选择，index 选择元素 elem
		(*used)[elem] = true
		*track = append(*track, elem)
		// 进入下一层回溯树
		backtrack(n, index+1, used, res, track)
		// 取消选择
		*track = (*track)[:len(*track)-1]
		(*used)[elem] = false
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func TestBeautifulArrangement(t *testing.T){
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