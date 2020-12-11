/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
package main
import "fmt"
func main() {
	a := [][]int{{1,2,3},{4,5,6},{7,8,9} }
	rotate(a)

	fmt.Printf("%v\n", a)

	b := [][]int{{5, 1, 9,11},{2, 4, 8,10}, {13, 3, 6, 7}, {15,14,12,16}}

	rotate(b)

	fmt.Printf("%v\n", b)
}

func rotate(matrix [][]int)  {

	n := len(matrix[0])

	res := [][]int{}

	for j := 0 ; j < n ; j++ {
		tmp := []int{}
		for i := n-1 ; i >= 0 ; i-- {
			tmp = append(tmp, matrix[i][j])
		}
		res = append(res, tmp)
	}
	for i := 0 ; i < n ; i ++ {
		for j := 0 ; j < n; j ++ {
			matrix[i][j] = res[i][j]
		}
	}
}
// @lc code=end

