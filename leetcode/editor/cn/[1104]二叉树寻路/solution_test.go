package leetcode

import (
	"math"
	"testing"
)

/**
在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。

 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；

 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。



 给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。



 示例 1：

 输入：label = 14
输出：[1,3,4,14]


 示例 2：

 输入：label = 26
输出：[1,2,6,10,26]




 提示：


 1 <= label <= 10^6


 Related Topics 树 数学 二叉树 👍 213 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

func pathInZigZagTree(label int) []int {
	path := []int{}
	for label >= 1 {
		path = append(path, label)
		label = label / 2
		depth := log(label)
		rangeArr := getLevelRange(depth)
		// 由于之字形分布，根据上层的节点取值范围，修正父节点
		//修正的过程其实是更新label这个循环变量的过程
		label = rangeArr[1] - (label - rangeArr[0])
	}
	// 反转成从根节点到目标节点的路径
	mid := len(path) / 2
	for i, j := 0, len(path)-1; i < mid; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// 获取第 n 层节点的取值范围
func getLevelRange(n int) []int {
	p := int(math.Pow(2, float64(n)))
	return []int{p, 2*p - 1}
}

func log(x int) int {
	return int(math.Log(float64(x)) / math.Log(2))
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathInZigzagLabelledBinaryTree(t *testing.T) {

}
