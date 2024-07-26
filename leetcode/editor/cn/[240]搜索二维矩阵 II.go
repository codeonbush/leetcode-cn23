package leetcode

import(
    "testing"
)

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性： 

 
 每行的元素从左到右升序排列。 
 每列的元素从上到下升序排列。 
 

 

 示例 1： 
 
 
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,2
3,26,30]], target = 5
输出：true
 

 示例 2： 
 
 
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,2
3,26,30]], target = 20
输出：false
 

 

 提示： 

 
 m == matrix.length 
 n == matrix[i].length 
 1 <= n, m <= 300 
 -10⁹ <= matrix[i][j] <= 10⁹ 
 每行的所有元素从左到右升序排列 
 每列的所有元素从上到下升序排列 
 -10⁹ <= target <= 10⁹ 
 

 Related Topics 数组 二分查找 分治 矩阵 👍 1495 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    // 初始化在右上角
    i, j := 0, n-1
    for i < m && j >= 0 {
        if matrix[i][j] == target {
            return true
        }
        if matrix[i][j] < target {
            // 需要大一点，往下移动
            i++
        } else {
            // 需要小一点，往左移动
            j--
        }
    }
    // while 循环中没有找到，则 target 不存在
    return false
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSearchA2dMatrixIi(t *testing.T){
    
}
