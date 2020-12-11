/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
package main

import "fmt"
import "strings"

func main() {
	var a string
	a = "/home/"
	fmt.Printf("%s, %s\n", a, simplifyPath(a))

	a = "/../"
	fmt.Printf("%s, %s\n", a, simplifyPath(a))


	a = "/home//foo/"
	fmt.Printf("%s, %s\n", a, simplifyPath(a))


	a = "/a/../../b/../c//.//"
	fmt.Printf("%s, %s\n", a, simplifyPath(a))

	a = "/a/./b/../../c/"
	fmt.Printf("%s, %s\n", a, simplifyPath(a))


	a = "/a//b////c/d//././/.."
	fmt.Printf("%s, %s\n", a, simplifyPath(a))
}

func simplifyPath(path string) string {
	stack := []string{}
	paths := strings.Split(path, "/")
	for i := 0 ; i < len(paths) ; i++ {
		if paths[i] == "" {
			continue
		}
		if paths[i] == "." {
			continue
		} else if paths[i] == ".." {
			tmp := len(stack) - 1
			if tmp < 0 {
				tmp = 0
			} 
			stack = stack[:tmp]
		} else {
			stack = append(stack, paths[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
// @lc code=end

