package leetcode

import (
	"testing"
)

/**
Given a 2D matrix matrix, handle multiple queries of the following type:


 Calculate the sum of the elements of matrix inside the rectangle defined by
its upper left corner (row1, col1) and lower right corner (row2, col2).


 Implement the NumMatrix class:


 NumMatrix(int[][] matrix) Initializes the object with the integer matrix
matrix.
 int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the
elements of matrix inside the rectangle defined by its upper left corner (row1,
col1) and lower right corner (row2, col2).


 You must design an algorithm where sumRegion works on O(1) time complexity.


 Example 1:


Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3,
 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Output
[null, 8, 11, 12]


Explanation
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0,
 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)



 Constraints:


 m == matrix.length
 n == matrix[i].length
 1 <= m, n <= 200
 -10⁴ <= matrix[i][j] <= 10⁴
 0 <= row1 <= row2 < m
 0 <= col1 <= col2 < n
 At most 10⁴ calls will be made to sumRegion.


 Related Topics 设计 数组 矩阵 前缀和 👍 566 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	// 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
	preSum [][]int
}

func Constructor_(matrix [][]int) NumMatrix {
	//m行 n列
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	// 构造前缀和矩阵
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum: preSum}
}

// 计算子矩阵 [x1, y1, x2, y2] 的元素和
func (this *NumMatrix) SumRegion(x1 int, y1 int, x2 int, y2 int) int {
	// 目标矩阵之和由四个相邻矩阵运算获得
	return this.preSum[x2+1][y2+1] - this.preSum[x1][y2+1] - this.preSum[x2+1][y1] + this.preSum[x1][y1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRangeSumQuery2dImmutable(t *testing.T) {

}
