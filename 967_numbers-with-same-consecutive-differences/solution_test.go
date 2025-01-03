package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。 

 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。 

 你可以按 任何顺序 返回答案。 

 

 示例 1： 

 
输入：n = 3, k = 7
输出：[181,292,707,818,929]
解释：注意，070 不是一个有效的数字，因为它有前导零。
 

 示例 2： 

 
输入：n = 2, k = 1
输出：[10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98] 

 示例 3： 

 
输入：n = 2, k = 0
输出：[11,22,33,44,55,66,77,88,99]
 

 示例 4： 

 
输入：n = 2, k = 2
输出：[13,20,24,31,35,42,46,53,57,64,68,75,79,86,97]
 

 

 提示： 

 
 2 <= n <= 9 
 0 <= k <= 9 
 

 Related Topics 广度优先搜索 回溯 👍 109 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func numsSameConsecDiff(n int, k int) []int {
	var res []int
	// 记录当前路径组成的数字的值
	track := 0
	// 记录当前数字的位数
	digit := 0
	backtrack(n, k, &res, &track, &digit)
	return res
}

// 回溯算法核心函数
func backtrack(n, k int, res *[]int, track *int, digit *int) {
	// base case，到达叶子节点
	if *digit == n {
		// 找到一个合法的 n 位数
		*res = append(*res, *track)
		return
	}

	// 回溯算法标准框架
	for i := 0; i <= 9; i++ {
		// 本题的剪枝逻辑 1，第一个数字不能是 0
		if *digit == 0 && i == 0 {
			continue
		}
		// 本题的剪枝逻辑 2，相邻两个数字的差的绝对值必须等于 k
		if *digit > 0 && abs(i-int(*track%10)) != k {
			continue
		}

		// 做选择，在 track 尾部追加数字 i
		*digit++
		*track = *track*10 + i
		// 进入下一层回溯树
		backtrack(n, k, res, track, digit)
		// 取消选择，删除 track 尾部数字
		*track /= 10
		*digit--
	}
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)


func TestNumbersWithSameConsecutiveDifferences(t *testing.T){
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