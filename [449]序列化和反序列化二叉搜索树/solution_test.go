package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

/**
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化
为最初的二叉搜索树。

 编码的字符串应尽可能紧凑。



 示例 1：


输入：root = [2,1,3]
输出：[2,1,3]


 示例 2：


输入：root = []
输出：[]




 提示：


 树中节点数范围是 [0, 10⁴]
 0 <= Node.val <= 10⁴
 题目数据 保证 输入的树是一棵二叉搜索树。


 Related Topics 树 深度优先搜索 广度优先搜索 设计 二叉搜索树 字符串 二叉树 👍 554 👎 0

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 分隔符，区分每个节点的值
const SEP = ","

// Encodes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	c.serializeHelper(root, &sb)
	result := sb.String()
	if len(result) > 0 && result[len(result)-1] == ',' {
		result = result[:len(result)-1]
	}
	return result
}

func (c *Codec) serializeHelper(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		return
	}
	// 前序遍历位置进行序列化
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)
	c.serializeHelper(root.Left, sb)
	c.serializeHelper(root.Right, sb)
}

// Decodes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	// 转化成前序遍历结果
	inorder := strings.Split(data, SEP)
	return c.deserializeHelper(&inorder, -0x3f3f3f3f, 0x3f3f3f3f)
}

// 定义：将 nodes 中值在闭区间 [min, max] 的节点构造成一棵 BST
func (c *Codec) deserializeHelper(nodes *[]string, min, max int) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	// 前序遍历位置进行反序列化
	// 前序遍历结果第一个节点是根节点
	rootVal, _ := strconv.Atoi((*nodes)[0])
	if rootVal > max || rootVal < min {
		// 超过闭区间 [min, max]，则返回空指针
		return nil
	}
	*nodes = (*nodes)[1:]
	// 生成 root 节点
	root := &TreeNode{
		Val: rootVal,
	}
	// 构建左右子树
	// BST 左子树都比根节点小，右子树都比根节点大
	root.Left = c.deserializeHelper(nodes, min, rootVal)
	root.Right = c.deserializeHelper(nodes, rootVal, max)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestSerializeAndDeserializeBst(t *testing.T) {

}
