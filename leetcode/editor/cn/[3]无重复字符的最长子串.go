package leetcode

import(
    "testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。 

 

 示例 1: 

 
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 

 示例 2: 

 
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 

 示例 3: 

 
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 

 

 提示： 

 
 0 <= s.length <= 5 * 10⁴ 
 s 由英文字母、数字、符号和空格组成 
 

 Related Topics 哈希表 字符串 滑动窗口 👍 10218 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
    window := make(map[byte]int)

    left, right := 0, 0
    res := 0 // 记录结果
    for right < len(s) {
        c := s[right]
        right++
        // 进行窗口内数据的一系列更新
        window[c]++
        // 判断左侧窗口是否要收缩
        for window[c] > 1 {
            d := s[left]
            left++
            // 进行窗口内数据的一系列更新
            window[d]--
        }
        // 在这里更新答案
        res = max(res, right - left)
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)


func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T){
    
}
