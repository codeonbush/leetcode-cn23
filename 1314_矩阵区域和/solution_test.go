package leetcode

import (
	"testing"
)

/**
给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[
r][c] 的和：


 i - k <= r <= i + k,
 j - k <= c <= j + k 且
 (r, c) 在矩阵内。




 示例 1：


输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
输出：[[12,21,16],[27,45,33],[24,39,28]]


 示例 2：


输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
输出：[[45,45,45],[45,45,45],[45,45,45]]




 提示：


 m == mat.length
 n == mat[i].length
 1 <= m, n, k <= 100
 1 <= mat[i][j] <= 100


 Related Topics 数组 矩阵 前缀和 👍 199 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	// 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	// 构造前缀和矩阵
	preSum := make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum}
}

// 计算子矩阵 [x1, y1, x2, y2] 的元素和
func (this *NumMatrix) SumRegion(x1, y1, x2, y2 int) int {
	// 目标矩阵之和由四个相邻矩阵运算获得
	return this.preSum[x2+1][y2+1] - this.preSum[x1][y2+1] - this.preSum[x2+1][y1] + this.preSum[x1][y1]
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	numMatrix := Constructor(mat)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 左上角的坐标
			x1 := max(i-k, 0)
			y1 := max(j-k, 0)
			// 右下角坐标
			x2 := min(i+k, m-1)
			y2 := min(j+k, n-1)

			res[i][j] = numMatrix.SumRegion(x1, y1, x2, y2)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMatrixBlockSum(t *testing.T) {

}
