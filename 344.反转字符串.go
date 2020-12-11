/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
package main
import "fmt"
import "strings"

func reverseString(s []byte)  {
	size := len(s)
	for i := 0 ; i < size/2 ; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
}
// @lc code=end

func main(){
  a := []byte("herlo")
  fmt.Println(string(a))
  reverseString(a)
  fmt.Println(string(a))

  b := []byte("h")
  fmt.Println(string(b))
  reverseString(b)
  fmt.Println(string(b))

  c := []byte("hai")
  fmt.Println(string(c))
  reverseString(c)
  fmt.Println(string(c))

   d := []byte("A man, a plan, a canal: Panama")
  fmt.Println(string(d))
  reverseString(d)
  fmt.Println(string(d))

  fmt.Println("result")
   aa := []string{"A"," ","m","a","n",","," ","a"," ","p","l","a","n",","," ","a"," ","c","a","n","a","l",":"," ","P","a","n","a","m","a"}
   fmt.Println(strings.Join(aa,""))
   fmt.Println(len(aa))

   aa2 := []string{"a","m","a","n","a","P"," ",":","l","a","n","a","c"," ","a"," ",",","n","a","l","p"," ","a"," ",",","n","a","m"," ","A"}
   fmt.Println(strings.Join(aa2,""))
}
