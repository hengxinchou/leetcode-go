/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
package main

// import "strings"
import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// for _, c := range ransomNote {
	// 	cc := string(c)
	// 	if strings.Contains(magazine, cc)  {
	// 		magazine = strings.Replace(magazine, cc, "", 1)
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true

	

	// if len(ransomNote) > len(magazine) {
	// 	return false
	// }

	// indexArray := make(map[rune]int)

	// for _, c := range magazine {
	// 	if _, ok := indexArray[c] ; ok {
	// 		indexArray[c] += 1
	// 	} else {
	// 		indexArray[c] = 1
	// 	}
	// }

	// for _, d := range ransomNote {
	// 	if value, ok := indexArray[d] ; ok {
	// 		tmp := value - 1
	// 		if tmp < 0 {
	// 			return false
	// 		} else {
	// 			indexArray[d] = tmp
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	// return true

	if len(ransomNote) > len(magazine) {
        return false
    }
 
    indexArray := map[rune]int{}

    for _, v := range ransomNote {
        index, _ := indexArray[v]
        index = strings.Index(magazine[index:], string(v))
        if index == -1 {
            return false
        }
        indexArray[v] += index + 1
    }
    
    return true
}
// @lc code=end

func main() {
	fmt.Printf("%t\n", canConstruct("aa", "ab"))
	fmt.Printf("%t\n", canConstruct("a", "b"))
	fmt.Printf("%t\n", canConstruct("aa", "aab"))
}

