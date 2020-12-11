/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

package main
import "fmt"
import "strings"
func main() {
	var s string
	var wordDict []string
	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	
	return
	s = "catsandogcat"
	wordDict = []string{"cats","dog","sand","and","cat","an"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	

	s = "ccbb"
	wordDict = []string{"bc", "cb"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	

	s =	"cars"
    wordDict = []string{"car","ca","rs"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	


	s = "leetcode"
	wordDict = []string{"leet", "code"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Printf("%s, %v, %t\n", s, wordDict, wordBreak(s, wordDict))	

}
 // @lc code=start
func wordBreak2(s string, wordDict []string) bool {
	
	var dfs func(string)

	n := len(wordDict)
	var isOk bool
	isOk = false
	dfs = func(left string) {
		if isOk == true {
			return
		}
		if len(left) == 0 {
			isOk = true
		}
		for  i := 0 ; i < n ; i ++ {
			index := strings.Index(left, wordDict[i]) 
			if index == 0 {
				newLeft := left[len(wordDict[i]):]
				dfs(newLeft)
			}
		}
	}
	dfs(s)

	return isOk
}
// @lc code=end


func wordBreak(s string, wordDict []string) bool {
    wordDictSet := make(map[string]bool)
    for _, w := range wordDict {
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordDictSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}

