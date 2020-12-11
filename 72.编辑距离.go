/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
package main
import (
	"fmt"
	"math"
	// "strings"
)

func main() {
	var a, b string

	a = "horse"	
	b = "ros"
	fmt.Printf("%s, %s, %d\n", a, b, minDistance(a,b))

	a = "intention"
	b = "execution"
	fmt.Printf("%s, %s, %d\n", a, b, minDistance(a,b))

	a = ""
	b = ""
	fmt.Printf("%s, %s, %d\n", a, b, minDistance(a,b))

	a = "zoologicoarchaeologist"
	b = "zoogeologist"
	fmt.Printf("%s, %s, %d\n", a, b, minDistance(a,b))
}
func minDistance(a string, b string) int {
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}

	n := len(a)
	m := len(b)

	states := make([][]int, n)
	for i := 0 ; i < n ; i++ {
		states[i] = make([]int, m)
	}

	for j := 0 ; j < m ; j++ {
		if a[0] == b[j] {
			states[0][j] = j
		} else if j != 0 {
			states[0][j] = states[0][j-1] + 1
		} else {
			states[0][j] = 1
		}
	}
	for i := 0 ; i < n ; i++ {
		if a[i] == b[0] {
			states[i][0] = i
		} else if i != 0 {
			states[i][0] = states[i-1][0] + 1
		} else {
			states[i][0] = 1
		}
	}

	for i := 1 ; i < n ; i++ {
		for j := 1 ; j < m ; j++ {

			leftUp := states[i-1][j-1]

			if a[i] != b[j]  {
				leftUp = states[i-1][j-1] + 1
			}
			states[i][j] = min( states[i-1][j] + 1, states[i][j-1] + 1, leftUp)
			// if a[i] == b[j] {
			// 	states[i][j] = min(states[i-1][j], states[i][j-1], states[i-1][j-1]) 
			// } else {
			// 	states[i][j] = min(states[i-1][j], states[i][j-1], states[i-1][j-1]) + 1
			// }
		}
	}
	// fmt.Println(strings.Split(b, ""))
	// for i := 0 ;i <n ; i ++ {
	// 		fmt.Printf("%s ", string(a[i]))
	// 		fmt.Println(states[i])
	// }
	// fmt.Println("world")
	return states[n-1][m-1]

}
// @lc code=end

func min (i, j, k int ) int {
	minV := math.MaxInt32

	if i < minV {
		minV = i
	} 
	if j < minV {
		minV = j
	}
	if k < minV {
		minV = k
	}
	return minV
}

