/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
package main

import "fmt"

func longestPalindrome(s string) int {
	a := map[byte]int{}
	for i := 0 ; i < len(s); i++{
		a[s[i]]++
	}
	sum := 0
	// maxSingle := 0
	flag := true // 第一个奇数使用，后面就不用
	fmt.Printf("size is %d\n", len(s))
	for k, v := range a {
		fmt.Printf("%s is %d\n", string(k), v)
		if v % 2 == 0 {
			sum += v
		} else if flag {
			sum += v
			flag = false
		} else {
			sum += v -1 
		}
	}
	return sum
}
// @lc code=end
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main(){
	a := "abccccdd"
	fmt.Println(longestPalindrome(a))

	a = ""
	fmt.Println(longestPalindrome(a))

	a = "x"
	fmt.Println(longestPalindrome(a))

	a = "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(a))
}
