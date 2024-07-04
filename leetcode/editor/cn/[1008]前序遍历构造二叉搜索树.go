package leetcode

import (
	"testing"
)

/**
给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。

 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。

 二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大于
Node.val。

 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。



 示例 1：




输入：preorder = [8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]


 示例 2:


输入: preorder = [1,3]
输出: [1,null,3]




 提示：


 1 <= preorder.length <= 100
 1 <= preorder[i] <= 10^8
 preorder 中的值 互不相同




 Related Topics 栈 树 二叉搜索树 数组 二叉树 单调栈 👍 290 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//前文 二叉树的花式构造 和 二叉树的序列化和反序列化 说过，生成二叉树的题目，无非就是
//先生成根节点，然后递归生成左右子树，最后把根节点和左右子树连接起来。具体区别在于你如何找到根节点，如何划分左右子树。
//
//根据前序遍历的特点是，根节点在第一位，后面接着左子树和右子树；BST 的特点，左子树都比根节点的值小，右子树都比根节点的值大。
//
//所以如何找到根节点？前序遍历的第一个就是根节点。
//
//如何找到左右子树？比根节点小的就是左子树的节点，比根节点大的就是右子树的节点。
//
//最后，确定清楚 build 函数的定义，利用这个函数递归生成 BST 即可。
func bstFromPreorder(preorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1)
}

// build函数返回根据 preorder[start..end] 区间内的元素生成的BST的根节点
func build(preorder []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	// 根据前序遍历的特点，根节点在第一位，后面接着左子树和右子树
	rootVal := preorder[start]
	root := &TreeNode{Val: rootVal}

	// 根据 BST 的特点，左子树都比根节点的值小，右子树都比根节点的值大
	// p 就是左右子树的分界点
	p := start + 1
	for p <= end && preorder[p] < rootVal {
		p++
	}
	// [start+1, p-1] 区间内是左子树元素
	root.Left = build(preorder, start+1, p-1)
	// [p, end] 区间内是右子树元素
	root.Right = build(preorder, p, end)

	return root
}

// gpt解法
func bstFromPreorder(preorder []int) *TreeNode {
	index := 0
	return buildBST(preorder, &index, int(^uint(0)>>1), -int(^uint(0)>>1)-1)
}

func buildBST(preorder []int, index *int, upper int, lower int) *TreeNode {
	if *index >= len(preorder) {
		return nil
	}

	val := preorder[*index]
	if val < lower || val > upper {
		return nil
	}

	*index++
	root := &TreeNode{Val: val}
	root.Left = buildBST(preorder, index, val, lower)
	root.Right = buildBST(preorder, index, upper, val)
	return root
}

//虽然逻辑正确，但是性能不够
//1. **频繁的切片操作**：
//   - 每次递归调用都会创建新的切片，例如 `preorder[1:rightIndex]` 和 `preorder[rightIndex:]`。虽然切片本身是对底层数组的引用，但每次切片操作都会创建新的切片头部结构，这会增加内存开销。
//   - 对于大规模数据，频繁的切片操作会导致大量的内存分配和复制，增加内存使用和运行时间。
//
//2. **递归深度**：
//   - 在最坏情况下（例如输入是一个严格递增或递减的序列），递归深度会达到数组的长度 `n`。每次递归调用都会在调用栈上分配内存，深度过大时会导致栈溢出。
//func bstFromPreorder(preorder []int) *TreeNode {
//	//第一个元素是根节点
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	rightIndex := 0
	for i := 0; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			rightIndex = i
			break
		}
	}

	if rightIndex > 0 {
		root.Left = bstFromPreorder(preorder[1:rightIndex])
	}
	root.Right = bstFromPreorder(preorder[rightIndex:])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinarySearchTreeFromPreorderTraversal(t *testing.T) {

}
