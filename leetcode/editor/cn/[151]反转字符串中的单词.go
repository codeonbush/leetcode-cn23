package leetcode

import(
    "strings"
    "testing"
)

/**
给你一个字符串 s ，请你反转字符串中 单词 的顺序。 

 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。 

 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。 

 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。 

 

 示例 1： 

 
输入：s = "the sky is blue"
输出："blue is sky the"
 

 示例 2： 

 
输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
 

 示例 3： 

 
输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
 

 

 提示： 

 
 1 <= s.length <= 10⁴ 
 s 包含英文大小写字母、数字和空格 ' ' 
 s 中 至少存在一个 单词 
 

 
 

 

 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。 

 Related Topics 双指针 字符串 👍 1167 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
    sb := strings.Builder{}
    // 先清洗一下数据，把多于的空格都删掉
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c != ' ' {
            // 单词中的字母/数字
            sb.WriteByte(c)
        } else if sb.Len() > 0 && sb.String()[sb.Len()-1] != ' ' {
            // 单词之间保留一个空格
            sb.WriteByte(' ')
        }
    }
    // 末尾如果有空格，清除之
    str := sb.String()
    if len(str) > 0 && str[len(str)-1] == ' ' {
        str = str[:len(str)-1]
    }

    // 清洗之后的字符串
    chars := []byte(str)
    n := len(chars)
    // 进行单词的翻转，先整体翻转
    reverse(chars, 0, n-1)
    // 再把每个单词翻转
    for i := 0; i < n; {
        for j := i; j < n; j++ {
            if j+1 == n || chars[j+1] == ' ' {
                // chars[i..j] 是一个单词，翻转之
                reverse(chars, i, j)
                // 把 i 置为下一个单词的首字母
                i = j + 2
                break
            }
        }
    }
    // 最后得到题目想要的结果
    return string(chars)
}

// 翻转 arr[i..j]
func reverse(arr []byte, i, j int) {
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}
//leetcode submit region end(Prohibit modification and deletion)


func TestReverseWordsInAString(t *testing.T){
    
}
