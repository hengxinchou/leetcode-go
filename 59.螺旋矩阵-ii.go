/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package main 
import "fmt"
// @lc code=start
func generateMatrix(n int) [][]int {
	a := make([][]int, n)
	for i := 0 ; i < n ; i++ {
		a[i] = make([]int, n)
	}


	k := 1
	l := 0
	r := n -1
	b := n -1
	t := 0
	det := n * n
	for k <= det {
		for i := t ; i <= r ; i++ {
			a[t][i] = k
			k++
		}
		t++ 
		for i := t ; i <= b ; i++ {
			a[i][r] = k
			k++
		}
		r-- 
		for i := r ; i >= l ; i-- {
			a[b][i] = k
			k++
		}
		b-- 
		for  i := b; i >= t ; i--{
			a[i][l] = k
			k++
		}
		l++
	}
	return a
}
// @lc code=end

func print(a [][]int) {
	n := len(a)
	for i := 0 ; i < n; i++{
		for j :=0 ; j< n; j++{
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}
func main () {
	result := 5 * 5
	fmt.Println("%d\n", result)
	a := generateMatrix(4)
	// fmt.Println(a)
	print(a)
	// var b int
	// b = result++
	// fmt.Println(b)
	// fmt.Println(result)
	// fmt.Println(resul t)

	// a := 0; b := 0
	// fmt.Println(a, b)
}