package leetcode

import(
    "testing"
)

/**
给你一个满足下述两条属性的 m x n 整数矩阵： 

 
 每行中的整数从左到右按非严格递增顺序排列。 
 每行的第一个整数大于前一行的最后一个整数。 
 

 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。 

 

 示例 1： 
 
 
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
 

 示例 2： 
 
 
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
 

 

 提示： 

 
 m == matrix.length 
 n == matrix[i].length 
 1 <= m, n <= 100 
 -10⁴ <= matrix[i][j], target <= 10⁴ 
 

 Related Topics 数组 二分查找 矩阵 👍 943 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    // 把二维数组映射到一维
    left, right := 0, m*n-1
    // 前文讲的标准的二分搜索框架
    for left <= right {
        mid := left + (right-left)/2
        if get(matrix, mid) == target {
            return true
        } else if get(matrix, mid) < target {
            left = mid + 1
        } else if get(matrix, mid) > target {
            right = mid - 1
        }
    }
    return false
}

// 通过一维坐标访问二维数组中的元素
func get(matrix [][]int, index int) int {
    n := len(matrix[0])
    // 计算二维中的横纵坐标
    i, j := index/n, index%n
    return matrix[i][j]
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSearchA2dMatrix(t *testing.T){
    
}
