package leetcode

import (
	"testing"
)

/**
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重
构得到原数据。

 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序
列化为原始的树结构。

 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法
解决这个问题。



 示例 1：


输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]


 示例 2：


输入：root = []
输出：[]


 示例 3：


输入：root = [1]
输出：[1]


 示例 4：


输入：root = [1,2]
输出：[1,2]




 提示：


 树中结点数在范围 [0, 10⁴] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 设计 字符串 二叉树 👍 1223 👎 0

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

//type Codec struct {
//
//}
//
//func Constructor() Codec {
//
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//
//}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 java 代码对比查看。

import (
	"strconv"
	"strings"
)

type Codec struct {
	SEP  string
	NULL string
}

func Constructor() Codec {
	return Codec{
		SEP:  ",",
		NULL: "#",
	}
}

/* 辅助函数，将二叉树存入 StringBuilder */
func (this *Codec) serialize(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(this.NULL)
		sb.WriteString(this.SEP)
		return
	}

	/******前序遍历位置******/
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(this.SEP)
	/***********************/

	this.serialize(root.Left, sb)
	this.serialize(root.Right, sb)
}

/* 主函数，将二叉树序列化为字符串 */
func (this *Codec) serializeTree(root *TreeNode) string {
	sb := &strings.Builder{}
	this.serialize(root, sb)
	return sb.String()
}

/* 辅助函数，通过 nodes 列表构造二叉树 */
func (this *Codec) deserialize(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	/******前序遍历位置******/
	// 列表最左侧就是根节点
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == this.NULL {
		return nil
	}
	root := &TreeNode{
		Val: atoi(first),
	}
	/***********************/

	root.Left = this.deserialize(nodes)
	root.Right = this.deserialize(nodes)

	return root
}

/* 主函数，将字符串反序列化为二叉树结构 */
func (this *Codec) deserializeTree(data string) *TreeNode {
	// 将字符串转化成列表
	nodes := strings.Split(data, this.SEP)
	return this.deserialize(&nodes)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestSerializeAndDeserializeBinaryTree(t *testing.T) {

}
