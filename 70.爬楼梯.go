/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
package main
import "fmt"
func climbStairs1(n int) int {
	p, q := 0, 0
	r := 1
	for i := 1; i <= n ; i++{
		p = q
		q = r
		r = p + q
		// p = q
		// p, q = q, p + q
	}
	return  r
	// return q
}

func climbStairs2(n int) int {
	if n == 2 {
		return 2
	} 
	if n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	ways := 0
	var dfs func(int) 
	map1 := map[int]bool{}
	dfs = func(i int) {
		if i == n {
			ways += 1
			return
		}
		// if _, ok := map1[i] ; ok {
		// 	return
		// } else{
		// 	map1[i] = true
		// }

		if i + 1 <= n {
		  dfs(i + 1)
		}

		if i + 2 <= n {
			dfs(i + 2)
		}


	}
	dfs(0)
	return ways
}
// @lc code=end

func main(){
	var a int
	a = 2
	fmt.Printf("%d, %d\n", a, climbStairs(a))
	
	a = 3
	fmt.Printf("%d, %d\n", a, climbStairs(a))

	a = 4
	fmt.Printf("%d, %d\n", a, climbStairs(a))

	a = 44
	fmt.Printf("%d, %d\n", a, climbStairs(a))
}