package leetcode

import (
	"testing"
)

/**
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



 示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]


 示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]




 提示：


 m == matrix.length
 n == matrix[i].length
 1 <= m, n <= 10
 -100 <= matrix[i][j] <= 100


 Related Topics 数组 矩阵 模拟 👍 1713 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper_bound, lower_bound := 0, m-1
	left_bound, right_bound := 0, n-1
	res := make([]int, 0, m*n)
	// len(res) == m * n 则遍历完整个数组
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			// 左边界右移
			left_bound++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSpiralMatrix(t *testing.T) {

}
