/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
package main
import "sort"
import "fmt"
import "strconv"

type stringArray []string

func (s stringArray) Len() int {
	return len(s)
}

func (s stringArray) Less(i, j int) bool {
	a := s[i]
	b := s[j]

	return a + b > b + a
	// if len(a) > len(b) {
	// 	dist := len(a)- len(b)
	// 	for i := 0 ;  i < dist ; i ++ {
	// 		b = b + string(b[0])
	// 	}
	// } else if len(b) > len(a) {
	// 	dist := len(b)- len(a)
	// 	for i := 0 ; i < dist ; i ++ {
	// 		a = a + string(a[0])
	// 	}
	// }
	
	// for i := 0 ; i < len(a) ; i++ {
	// 	if a[i] > b[i] {
	// 		return true
	// 	} else if a[i] < b[i] {
	// 		return false
	// 	}
	// }

	// c := s[i]
	// d := s[j]
	// if c[len(c)-1] > d[len(d)-1] {
	// 	return true
	// } else {
	// 	return false
	// }

	// fmt.Printf("a is %s, b is %s, len is %d, flag is %t\n", a, b, len(a), flag)
	// return flag

}

func (s stringArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	a := make([]string, len(nums))
	for i, v := range nums {
		a[i] = strconv.Itoa(v)
	}
	// sort.Slice()
	sort.Sort(stringArray(a))
	fmt.Printf("string Array is %v\n", a)
	if a[0] == "0" {
		return "0"
	}

	b := ""
	for _, v := range a {

		b += v
	}
	return b
}
// @lc code=end

func main (){
	a1 := []int{10, 2}
	fmt.Printf("%v, largestNumber is %s\n", a1, largestNumber(a1))
	a2 := []int{3, 30,34,5,9}
	fmt.Printf("%v, largestNumber is %s\n", a2, largestNumber(a2))
	a3 := []int{128, 12}
	fmt.Printf("%v, largestNumber is %s\n", a3, largestNumber(a3))

	a4 := []int{12, 121}
	fmt.Printf("%v, largestNumber is %s\n", a4, largestNumber(a4))

	a5 := []int{20, 1}
	fmt.Printf("%v, largestNumber is %s\n", a5, largestNumber(a5))
	fmt.Printf("hello\n")

	a6 := []int{0, 0}
	fmt.Printf("%v, largestNumber is %s\n", a6, largestNumber(a6))
	fmt.Printf("hello\n")
}