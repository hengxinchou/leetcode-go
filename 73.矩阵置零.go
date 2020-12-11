package main
import "fmt"
func main() {
	a := [][]int{{1,0,1},{1,0,1},{1,1,1}}
	print(a)
	setZeroes(a)
	print(a)
}

func print(a [][]int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func print2(a [][]bool) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	
	visited := make([][]bool, m)
	for i := 0 ; i < m ; i++ {
		visited[i] = make([]bool, n)
	}
	// print2(visited)
	// fmt.Printf("m is %d, n is %d\n", m, n)
	for i := 0 ; i < m ; i++ {
		for j := 0 ; j < n ; j++ {
			// fmt.Printf("i is %d, j is %d\n", i, j)
			// if visited[i][j] {
				// continue
			// }
			if !visited[i][j]  && matrix[i][j] == 0 {
				// fmt.Printf("zero i is %d, j is %d\n", i, j)
				visited[i][j] = true
				for k := 0 ; k < n ; k++ {
					if visited[i][k] || matrix[i][k] == 0 {
						continue
					}
					matrix[i][k] = 0 
					visited[i][k] = true
					// fmt.Printf("mark visited[%d][%d]\n", i, k)
				}
				// fmt.Println("hello")
				for k := 0 ; k < m ; k++ {
					if visited[k][j] || matrix[k][j] == 0 {
						continue
					}
					matrix[k][j] = 0 
					visited[k][j] = true
					// fmt.Printf("mark visited[%d][%d]\n",k, j)
				}
			}
		}
	}
}
// @lc code=end

