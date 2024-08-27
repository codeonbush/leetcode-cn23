package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。

 当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。


 例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。


 实现 StockSpanner 类：


 StockSpanner() 初始化类对象。
 int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。




 示例：


输入：
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
输出：
[null, 1, 1, 1, 2, 1, 4, 6]

解释：
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // 返回 1
stockSpanner.next(80);  // 返回 1
stockSpanner.next(60);  // 返回 1
stockSpanner.next(70);  // 返回 2
stockSpanner.next(60);  // 返回 1
stockSpanner.next(75);  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
stockSpanner.next(85);  // 返回 6




 提示：


 1 <= price <= 10⁵
 最多调用 next 方法 10⁴ 次


 Related Topics 栈 设计 数据流 单调栈 👍 455 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type StockSpanner struct {
	// int[] 记录 {价格，小于等于该价格的天数} 二元组
	stk [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{stk: make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	// 算上当天
	count := 1
	// 单调栈模板
	for len(this.stk) > 0 && price >= this.stk[len(this.stk)-1][0] {
		// 挤掉价格低于 price 的记录
		prev := this.stk[len(this.stk)-1]
		this.stk = this.stk[:len(this.stk)-1]
		// 计算小于等于 price 的天数
		count += prev[1]
	}
	this.stk = append(this.stk, [2]int{price, count})

	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestOnlineStockSpan(t *testing.T) {
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
			result := Next(test.input)
			fmt.Println(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For heights %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
