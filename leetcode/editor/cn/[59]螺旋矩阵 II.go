package leetcode

import(
    "testing"
)

/**
给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。 

 

 示例 1： 
 
 
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
 

 示例 2： 

 
输入：n = 1
输出：[[1]]
 

 

 提示： 

 
 1 <= n <= 20 
 

 Related Topics 数组 矩阵 模拟 👍 1301 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    upper_bound, lower_bound := 0, n-1
    left_bound, right_bound := 0, n-1
    // 需要填入矩阵的数字
    num := 1

    for num <= n*n {
        if upper_bound <= lower_bound {
            // 在顶部从左向右遍历
            for j := left_bound; j <= right_bound; j++ {
                matrix[upper_bound][j] = num
                num++
            }
            // 上边界下移
            upper_bound++
        }

        if left_bound <= right_bound {
            // 在右侧从上向下遍历
            for i := upper_bound; i <= lower_bound; i++ {
                matrix[i][right_bound] = num
                num++
            }
            // 右边界左移
            right_bound--
        }

        if upper_bound <= lower_bound {
            // 在底部从右向左遍历
            for j := right_bound; j >= left_bound; j-- {
                matrix[lower_bound][j] = num
                num++
            }
            // 下边界上移
            lower_bound--
        }

        if left_bound <= right_bound {
            // 在左侧从下向上遍历
            for i := lower_bound; i >= upper_bound; i-- {
                matrix[i][left_bound] = num
                num++
            }
            // 左边界右移
            left_bound++
        }
    }
    return matrix
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSpiralMatrixIi(t *testing.T){
    
}
