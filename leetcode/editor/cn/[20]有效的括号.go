package leetcode

import(
    "testing"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 

 有效字符串需满足： 

 
 左括号必须用相同类型的右括号闭合。 
 左括号必须以正确的顺序闭合。 
 每个右括号都有一个对应的相同类型的左括号。 
 

 

 示例 1： 

 
输入：s = "()"
输出：true
 

 示例 2： 

 
输入：s = "()[]{}"
输出：true
 

 示例 3： 

 
输入：s = "(]"
输出：false
 

 

 提示： 

 
 1 <= s.length <= 10⁴ 
 s 仅由括号 '()[]{}' 组成 
 

 Related Topics 栈 字符串 👍 4507 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(str string) bool {
    pair := map[rune]rune{')': '(', '}': '{', ']': '['}
    var left stack

    for _, c := range str {
        if isOpening(c) {
            left.push(c)
        } else if matching, found := pair[c]; found {
            if left.isEmpty() || left.pop() != matching {
                return false
            }
        } else {
            return false // 无效的右括号
        }
    }

    return left.isEmpty()
}

func isOpening(c rune) bool {
    return c == '(' || c == '{' || c == '['
}

// stack implementation remains the same
type stack []rune

func (s *stack) push(v rune) {
    *s = append(*s, v)
}

func (s *stack) pop() rune {
    if s.isEmpty() {
        return ' ' // Handle underflow if needed
    }
    old := *s
    n := len(old)
    v := old[n-1]
    *s = old[0 : n-1]
    return v
}

func (s *stack) peek() rune {
    if s.isEmpty() {
        return ' ' // Handle underflow if necessary
    }
    return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
    return len(*s) == 0
}
//leetcode submit region end(Prohibit modification and deletion)


func TestValidParentheses(t *testing.T){
    
}
