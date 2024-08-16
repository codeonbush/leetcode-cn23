package leetcode

import (
	"container/list"
	"strings"
	"testing"
	"unicode/utf8"
)

/**
假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：



 这里将 dir 作为根目录中的唯一目录。dir 包含两个子目录 subdir1 和 subdir2 。subdir1 包含文件 file1.ext 和子目录
subsubdir1；subdir2 包含子目录 subsubdir2，该子目录下包含文件 file2.ext 。

 在文本格式中，如下所示(⟶表示制表符)：


dir
⟶ subdir1
⟶ ⟶ file1.ext
⟶ ⟶ subsubdir1
⟶ subdir2
⟶ ⟶ subsubdir2
⟶ ⟶ ⟶ file2.ext


 如果是代码表示，上面的文件系统可以写为 "dir
\tsubdir1
\t\tfile1.ext
\t\tsubsubdir1
\tsubdir2
\t\tsubsubdir2
\t\t\tfile2.ext" 。'
' 和 '\t' 分别是换行符和制表符。

 文件系统中的每个文件和文件夹都有一个唯一的 绝对路径 ，即必须打开才能到达文件/目录所在位置的目录顺序，所有路径用 '/' 连接。上面例子中，指向 file2
.ext 的 绝对路径 是 "dir/subdir2/subsubdir2/file2.ext" 。每个目录名由字母、数字和/或空格组成，每个文件名遵循
name.extension 的格式，其中
 name 和
 extension由字母、数字和/或空格组成。

 给定一个以上述格式表示文件系统的字符串 input ，返回文件系统中 指向 文件 的 最长绝对路径 的长度 。 如果系统中没有文件，返回 0。



 示例 1：


输入：input = "dir
\tsubdir1
\tsubdir2
\t\tfile.ext"
输出：20
解释：只有一个文件，绝对路径为 "dir/subdir2/file.ext" ，路径长度 20


 示例 2：


输入：input = "dir
\tsubdir1
\t\tfile1.ext
\t\tsubsubdir1
\tsubdir2
\t\tsubsubdir2
\t\t\tfile2.ext"
输出：32
解释：存在两个文件：
"dir/subdir1/file1.ext" ，路径长度 21
"dir/subdir2/subsubdir2/file2.ext" ，路径长度 32
返回 32 ，因为这是最长的路径

 示例 3：


输入：input = "a"
输出：0
解释：不存在任何文件

 示例 4：


输入：input = "file1.txt
file2.txt
longfile.txt"
输出：12
解释：根目录下有 3 个文件。
因为根目录中任何东西的绝对路径只是名称本身，所以答案是 "longfile.txt" ，路径长度为 12




 提示：


 1 <= input.length <= 10⁴
 input 可能包含小写或大写的英文字母，一个换行符 '
'，一个制表符 '\t'，一个点 '.'，一个空格 ' '，和数字。


 Related Topics 栈 深度优先搜索 字符串 👍 281 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func lengthLongestPath(input string) int {
	// 栈保存各级目录的长度
	stack := list.New()
	maxLen := 0
	// 按行分割输入，使用 \ 作为分隔符
	parts := strings.Split(input, "\n")

	for _, part := range parts {
		// 计算当前目录的层级
		level := 0
		for level < len(part) && part[level] == '\t' {
			level++
		}

		// 确保栈中只保留当前目录的父路径
		for stack.Len() > level {
			stack.Remove(stack.Back())
		}

		// 获取当前目录或文件的名称
		current := part[level:]
		// 更新当前路径长度
		currentLen := utf8.RuneCountInString(current)

		// 如果是文件，计算当前路径的总长度
		if strings.Contains(current, ".") {
			sum := 0
			for e := stack.Front(); e != nil; e = e.Next() {
				sum += utf8.RuneCountInString(e.Value.(string))
			}
			// 加上父路径的分隔符
			sum += currentLen + stack.Len() // + stack.Len() 是路径分隔符的数量
			maxLen = max(maxLen, sum)
		} else {
			// 如果是目录，添加到栈中
			stack.PushBack(current)
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestAbsoluteFilePath(t *testing.T) {

}
