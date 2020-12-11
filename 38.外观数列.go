/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
package main
import "fmt"
import "strconv"
import "strings"

func countAndSay(n int) string {
	// size := 30
	// a := make([]string, size)
	// cur := "1"
	prev := "1"
	for i := 1 ; i < n ; i++ {
		j := 0
		var cur strings.Builder
		for j < len(prev)  {
			v := prev[j]
			count := 0
			for  j < len(prev) && prev[j] == v {
				count++
				j++
			}
			cur.WriteString(strconv.Itoa(count))
			cur.WriteByte(v)
		}
		// fmt.Println(cur)
		prev = cur.String()
	}
	return prev
	// res := make([]string, 30)
	// for i := 0 ; i < 30 ; i++ {
	// 	str := ""
	// 	for _, s := range a[i] {
	// 		str += s
	// 	}
	// 	fmt.Println(str)
	// 	res[i] = str
	// }
	// return res[n]
}
// @lc code=end

func main(){
	// b := countAndSay(30)
	for i := 1 ; i < 30 ; i++ {
		fmt.Println(countAndSay(i))
	}
}

