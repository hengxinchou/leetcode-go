/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
package main

func main() {
	a := [][]byte{
		{'A','B', 'C', 'E'},
		{'S', 'F','C', 'S'},
		{'A', 'D', 'E', 'E'}
	}
	for i := 0 ;i < len(a) ; i++ {
		fmt.Println(a[i])
	}
	var word string

	word = "ABCCED"
	fmt.Printf("%s, %t\n", word, exist(a, word))

	word = "SEE"
	fmt.Printf("%s, %t\n", word, exist(a, word))

	word = "ABCB"
	fmt.Printf("%s, %t\n", word, exist(a, word))
}
func exist(board [][]byte, word string) bool {
	match := false
	var dfs(int, int, int)
	dfs = func(row, column, index int) {
		if world[index] == board[row][column] && index == len(word) - 1 {
			match = true
			return
		}
		if board[row][column] == word[index] {
			dfs(row, column + 1, 1)

		} else {
			dfs(1,1,0)
		}
	}
	dfs(0, 0, 0)
	return match
}

// @lc code=end

