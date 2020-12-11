/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start

package main

import "fmt"
import "sort"
// import "strings"

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	b := groupAnagrams(a)
	fmt.Printf("%v, %v\n", a, b)
}


type bytes []byte

func (b bytes) Len() int {
    return len(b)
}
func (b bytes) Less(i,j int) bool {
    return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}



func transfer(s string) string {
	// res := []string{}
	// for i := 0 ; i< len(s) ; i++ {
	// 	res = append(res, string(s[i]))
	// }
	// sort.Strings(res)

	res := bytes(s)
	sort.Sort(res)

	return string(res)
}

func groupAnagrams(strs []string) [][]string {
	mapStrings := map[string][]string{}
	for _, s := range strs {
		transferS := transfer(s)
		if stringsSlice, ok := mapStrings[transferS] ; ok {
			mapStrings[transferS] = append(stringsSlice, s)
		} else {
			mapStrings[transferS] = []string{s}
		}
	}
	res := [][]string{}
	for _, v := range mapStrings {
		res = append(res, v)
	}
	return res
}

// @lc code=end

