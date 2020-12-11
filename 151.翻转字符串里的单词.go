/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
package main
import "strings"
import "fmt"
func removeBlank(s string) string {
	var s2 strings.Builder
	l := len(s)
	for s[l-1] == ' ' {
		l--
	}
	for i := 0; i < l; i++ {
		if i == 0  && s[i] == ' ' {
			continue
		} 
		if i > 0 && s[i] == ' ' && s[i-1] == ' ' {
			continue
		}	
		s2.WriteByte(s[i])
	}
	return s2.String()

}
func reverseWords(s string) string {

	// fmt.Println(s2.String())
	s2 := removeBlank(s)
	s3 := []byte(s2)

	size := len(s3)
	// fmt.Printf("size is %d\n", size)
	for i := 0 ; i < size / 2 ; i++ {
		s3[i], s3[size -i -1] = s3[size -i -1], s3[i]
	}
	// fmt.Printf("inverse s3 is %s\n", string(s3))
	// fmt.Printf("size is %d\n", size)
	i := 0
	for i < size {
		j := 0
		for i + j < size && s3[i+j] != ' ' {
			j++
		}
		start := i 
		end := start + j 
		// fmt.Printf("start is %d, end is %d, i is %d, j is %d\n", start ,end , i,j)
		for  k := 0 ; k < j  / 2 ; k++ {
			// fmt.Printf("start + i is %d, end-i-1 is %d\n", start +i , end -i -1)
			s3[start + k], s3[end-k-1] = s3[end-k-1], s3[start+k]
		}
		i = end + 1
	}
	// s2 := removeBlank(string(s3))
	// return removeBlank(string(s3))
	return string(s3)
}
// @lc code=end

func main(){
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("  Bob    Loves  Alice "))
	fmt.Println(reverseWords("Alice does not even like bob"))
}

