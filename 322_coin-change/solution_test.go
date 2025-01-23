package leetcode

import(
	"fmt"
	"math"
	"reflect"
	"testing"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。 

 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。 

 你可以认为每种硬币的数量是无限的。 

 

 示例 1： 

 
输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1 

 示例 2： 

 
输入：coins = [2], amount = 3
输出：-1 

 示例 3： 

 
输入：coins = [1], amount = 0
输出：0
 

 

 提示： 

 
 1 <= coins.length <= 12 
 1 <= coins[i] <= 2³¹ - 1 
 0 <= amount <= 10⁴ 
 

 Related Topics 广度优先搜索 数组 动态规划 👍 2964 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount + 1)
	for i := range memo {
		memo[i] = -666
	}
	// 备忘录初始化为一个不会被取到的特殊值，代表还未被计算
	return dp(coins, amount, memo)
}

func dp(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 查备忘录，防止重复计算
	if memo[amount] != -666 {
		return memo[amount]
	}

	res := math.MaxInt32
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dp(coins, amount - coin, memo)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		res = min(res, subProblem + 1)
	}
	// 把计算结果存入备忘录
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)


func TestCoinChange(t *testing.T){
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