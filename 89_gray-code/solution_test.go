package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
n 位格雷码序列 是一个由 2ⁿ 个整数组成的序列，其中：

 
 每个整数都在范围 [0, 2ⁿ - 1] 内（含 0 和 2ⁿ - 1） 
 第一个整数是 0 
 一个整数在序列中出现 不超过一次 
 每对 相邻 整数的二进制表示 恰好一位不同 ，且 
 第一个 和 最后一个 整数的二进制表示 恰好一位不同 
 

 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。 

 

 示例 1： 

 
输入：n = 2
输出：[0,1,3,2]
解释：
[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
- 00 和 01 有一位不同
- 01 和 11 有一位不同
- 11 和 10 有一位不同
- 10 和 00 有一位不同
[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
- 00 和 10 有一位不同
- 10 和 11 有一位不同
- 11 和 01 有一位不同
- 01 和 00 有一位不同
 

 示例 2： 

 
输入：n = 1
输出：[0,1]
 

 

 提示： 

 
 1 <= n <= 16 
 

 Related Topics 位运算 数学 回溯 👍 693 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func grayCode(n int) []int {
	used := make(map[int]struct{})
	path := []int{}
	var res []int
	traverse(0, n, &used, &path, &res)
	return res
}

func traverse(root, n int, used *map[int]struct{}, path *[]int, res *[]int) {
	if *res != nil {
		return
	}

	if len(*path) == (1 << n) {
		*res = append([]int(nil), *path...)
		return
	}

	if _, found := (*used)[root]; found {
		return
	}

	// 多叉树遍历的前序位置
	(*used)[root] = struct{}{}
	*path = append(*path, root)

	// 对当前数字的每个二进制位进行翻转，得到子节点
	for i := 0; i < n; i++ {
		next := flipBit(root, i)
		traverse(next, n, used, path, res)
	}

	// 多叉树遍历的后序位置
	delete(*used, root)
	*path = (*path)[:len(*path)-1]
}

// 把第 i 位取反（0 变 1，1 变 0）
func flipBit(x, i int) int {
	return x ^ (1 << i)
}
//leetcode submit region end(Prohibit modification and deletion)


func TestGrayCode(t *testing.T){
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