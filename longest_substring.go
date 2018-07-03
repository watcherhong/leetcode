package leetcode

import (
    "fmt"
)

// func main() {
//     fmt.Println("Longest Substring")
// }

// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// Given "ebfcadbcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
    var longest int = 0
    str := []byte(s)
    m := make(map[byte]int)
    valid := 0
    for k, v := range str {
        if i, ok := m[v]; ok {
            if valid <= i {
                if l := k - valid; l > longest {
                    longest = l
                }
                valid = i + 1
            }
        }
        m[v] = k
    } 
    //fmt.Printf("......longest=%d valid=%d len(%s)=%d \n", longest, valid, str, len(str))
    if last := len(str) - valid; last > longest {
        longest = last
    }
    return longest
}

func TestLongestSubstring() {
    fmt.Println("TestLongestSubstring")
    var str string
    str = "abcabcbb"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "bbbbb"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "c"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "aux"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "ebfcadbcbb"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "aab"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
    str = "abba"
    fmt.Printf("The longest substring of %s is %d \n", str, lengthOfLongestSubstring(str))
}