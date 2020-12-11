/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var version1, version2 string

	version1 = "1.01"
	version2 = "1.001"
	fmt.Printf("%s, %s, %d\n", version1, version2, compareVersion(version1, version2))

	version1 = "1.0"
	version2 = "1.0.0"
	fmt.Printf("%s, %s, %d\n", version1, version2, compareVersion(version1, version2))

	version1 = "0.1"
	version2 = "1.1"
	fmt.Printf("%s, %s, %d\n", version1, version2, compareVersion(version1, version2))

	version1 = "1.0.1"
	version2 = "1"
	fmt.Printf("%s, %s, %d\n", version1, version2, compareVersion(version1, version2))

	version1 = "7.5.2.4"
	version2 = "7.5.3"
	fmt.Printf("%s, %s, %d\n", version1, version2, compareVersion(version1, version2))
}

func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")

	var tmp1, tmp2 int
	var i int
	for i = 0; i < len(arr1) && i < len(arr2); i++ {
		tmp1, _ = strconv.Atoi(arr1[i])
		tmp2, _ = strconv.Atoi(arr2[i])
		if tmp1 > tmp2 {
			return 1
		} else if tmp1 == tmp2 {
			continue
		} else {
			return -1
		}
	}

	for ; i < len(arr1); i++ {
		tmp1, _ = strconv.Atoi(arr1[i])
		if tmp1 > 0 {
			return 1
		}
		if tmp1 == 0 {
			if i == len(arr1)-1 {
				return 0
			} else {
				continue
			}
		}
	}

	for ; i < len(arr2); i++ {
		tmp2, _ = strconv.Atoi(arr2[i])
		if tmp2 > 0 {
			return -1
		}
		if tmp2 == 0 {
			if i == len(arr2)-1 {
				return 0
			} else {
				continue
			}
		}
	}
	return 0
}

// @lc code=end
