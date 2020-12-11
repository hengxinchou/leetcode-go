/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
package main
import "fmt"
func main() {
	var a []int	
	a = []int{4,2,5,7}
	fmt.Println(a)
	fmt.Printf("%v\n", sortArrayByParityII(a))

	a = []int{648,831,560,986,192,424,997,829,897,843}
	fmt.Println(a)
	fmt.Printf("%v\n", sortArrayByParityII(a))
}

func sortArrayByParityII(A []int) []int {
	// var odd, even int
	if len(A) < 1 {
		return A
	}

	var tmp int
	queueOdd := []int{}
	queueEven := []int{}
	
	for i := 0 ; i < len(A) ; i++ {
		if i % 2 == 0 && A[i] % 2 != 0 {
			if len(queueOdd) > 0 {
				// fmt.Printf("even: %d, %d, %d, %d\n", odd, A[odd], even, A[even])
				tmp = queueOdd[0]
				queueOdd = queueOdd[1:]
				A[i], A[tmp] = A[tmp], A[i]
			} else {
				queueEven = append(queueEven, i)
			}
		} else if i % 2 == 1 && A[i] % 2 != 1 {
			if len(queueEven) > 0 {
				// fmt.Printf("odd: %d, %d, %d, %d\n", odd, A[odd], even, A[even])
				tmp = queueEven[0]
				queueEven = queueEven[1:]
				A[i], A[tmp] = A[tmp], A[i]
			} else {
				queueOdd = append(queueOdd, i)
			}
		}
	}
	return A
}
// @lc code=end

