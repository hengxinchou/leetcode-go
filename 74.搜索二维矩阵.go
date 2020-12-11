/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
package main
import "fmt"

func main() {
	var a [][]int
	a = [][]int{
	   {1,3,5,7},
	   {10,11,16,20},
	   {23,30,34,50},
	}	
	for i := 0 ; i < len(a) ; i++ {
		fmt.Println(a[i])
	}
	var target  int

	target = 3
	fmt.Printf("%d, %t\n", target,  searchMatrix(a, target))

	target = 1
	fmt.Printf("%d, %t\n", target,  searchMatrix(a, target))

	target = 23
	fmt.Printf("%d, %t\n", target,  searchMatrix(a, target))

	target = 50
	fmt.Printf("%d, %t\n", target,  searchMatrix(a, target))

	target = 30
	fmt.Printf("%d, %t\n", target,  searchMatrix(a, target))

	target = 33
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))
	
	target = -1
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))

	a = [][]int{}
	target = -1
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))

	a = [][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,50},
	}
	for i := 0 ; i < len(a) ; i++ {
		fmt.Println(a[i])
	}
	target = 13
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))

	a = [][]int{{}}
	target = 1
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))

	a = [][]int{
		{1}, {3},
	}
	for i := 0 ; i < len(a) ; i++ {
		fmt.Println(a[i])
	}
	target = 3
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))
	target = 1
	fmt.Printf("%d, %t\n",  target,  searchMatrix(a, target))
}

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	row := len(matrix) 
	column := len(matrix[0]) 

	if column <= 0 {
		return false
	}

	low := 0 
	upper := row - 1
	var mid int
	targetRow  := -1
	for low <= upper {
		// fmt.Printf("%d, %d\n", low, upper)
		if low < 0 || upper > row {
			break
		}
		mid = low + ((upper - low) >> 1 )
		// fmt.Printf("mid is %d\n", mid)
		if matrix[mid][0] <= target && matrix[mid][column -1 ] >= target && ( mid == row - 1 || matrix[mid+1][0] >= target ) {
			targetRow = mid
			break
		}
		if matrix[mid][0] <= target {
			low = mid + 1
		} else {
			upper = mid  - 1
		}
	}

	// fmt.Printf("target row is %d\n", targetRow)
	if targetRow == -1 {
		return false
	}
	i := targetRow
	
	low = 0
	upper = column - 1
	for low <= upper {
		mid = low + (( upper - low) >> 1)

		if matrix[i][mid] == target {
			return true
		}

		if matrix[i][mid] < target {
			low = mid + 1
		} else {
			upper = mid - 1
		}
	} 
	return false
}
// @lc code=end


func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
    if m == 0 {
		return false
	} 
    n := len(matrix[0])

    // 二分查找
	low := 0
	upper := m * n - 1
    var pivotIdx, pivotElement, row, column int
    for low <= upper {
	  pivotIdx = low + (upper - low) >> 1
	  row  = pivotIdx / n
	  column = pivotIdx % n
      pivotElement = matrix[row][column]
      if target == pivotElement {
		  return true
	   }
	   if target < pivotElement {
	    	upper = pivotIdx - 1
	   } else { 
		    low = pivotIdx + 1
	   }
    }
    return false
}