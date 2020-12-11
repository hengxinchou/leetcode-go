/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start
// package main 
// import "fmt"
// import "sort"
// // import "reflect"

// type entry struct {
// 	value []int
// 	sum int
// }

// type entryArray []*entry

// func (ea entryArray) Len() int {
// 	return len(ea)
// }

// func (ea entryArray)Less(i, j int) bool {
// 	return ea[i].sum < ea[j].sum
// }

// func (ea entryArray) Swap(i, j int) {
// 	ea[i].value, ea[j].value = ea[j].value, ea[i].value
// 	ea[i].sum, ea[j].sum = ea[j].sum, ea[i].sum
// }

// func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
// 	ea := make(entryArray, len(nums1) * len(nums2))
// 	// ea := []*entry{}
// 	l := 0
// 	for _, i := range nums1 {
// 		for _, j := range nums2 {
// 			// sum = i + j
// 			ea[l] = &entry{value: []int{i, j}, sum: (i + j)}
// 			l++
// 		}
// 	}
// 	// printArray(ea)
// 	// fmt.Println("after sort")
// 	// fmt.Printf("%T\n", ea)
// 	// fmt.Printf("%s\n", reflect.TypeOf(ea))
// 	sort.Sort(&ea)
// 	// printArray(ea)

// 	a := [][]int{}

// 	m := 0
// 	for m < k && m < len(ea) {
// 		a = append(a, ea[m].value)
// 		m++
// 	}
// 	return a
// }
// // @lc code=end

// func printArray(a entryArray) {
// 	for _, i := range a {
// 		fmt.Printf("%+v\n", i.value)
// 	}
// 	fmt.Println()
	
// }
// func print(nums[][]int){
// 	for _, x := range nums {
// 		fmt.Printf("is %v\n", x)
// 	}
// 	fmt.Println()
// }


// func main(){
// 	nums1 := []int{1,2}
// 	nums2 := []int{3}
// 	// var i = 4
// 	// fmt.Println(i)
// 	result := kSmallestPairs(nums1, nums2, 3)
// 	fmt.Printf("result is %v\n", result)
// 	print(result)
// }

package main
import "fmt"
import "container/heap"

type entry struct {
	i int
	j int 
	priority int
}
type Pq []*entry

func (pq Pq) Len() int{
	return len(pq)
}

func (pq Pq) Swap(i, j int) {
	pq[i].i, pq[i].j, pq[j].i, pq[j].j = pq[j].i, pq[j].j, pq[i].i, pq[i].j
	// pq[i], pq[j] = pq[j], pq[i]
	pq[i].priority, pq[j].priority = pq[j].priority, pq[i].priority
}
func (pq Pq) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *Pq) Push(x interface{}) {
	entry := x.(*entry)
	*pq = append(*pq, entry)
}

func (pq *Pq) Pop() interface{} {
	n := pq.Len()
	entry := (*pq)[n-1]
	*pq = (*pq)[0:n-1]
	return entry
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// pq := make(Pq, len(nums1) * len(nums2))
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	l1 := len(nums1)
	l2 := len(nums2)

	pq := new(Pq)
	heap.Init(pq)
	// for _, i := range nums1 {
	// 	for _, j := range nums2 {
	// 		// fmt.Printf("i is %d, j is %d\n", i, j)
	// 	 	heap.Push(pq, &entry{i: i, j:j , priority: i + j})
	// 	}
	// }
	heap.Push(pq, &entry{i: 0, j: 0, priority:nums1[0] + nums2[0]})
	// print(pq)

	visited := make([][]bool, l1)
	for k, _ := range visited{
		visited[k] = make([]bool, l2)
	}

	res := make([][]int, 0, k)
	
	for pq.Len() > 0  &&  len(res) < k {
		a := heap.Pop(pq).(*entry)
		i := a.i
		j := a.j

		if visited[i][j] {
			continue
		}

		// fmt.Printf("pop %d, %d, %d\n", a.i, a.j, a.priority)
		// print(pq)
		// fmt.Println()
		res  = append(res, []int{nums1[i], nums2[j]})
		visited[i][j]=true
		if i + 1 < l1 {
			heap.Push(pq,  &entry{i: i+1, j: j, priority: nums1[i+1] + nums2[j]})
		} 
		if j + 1 < l2 {
			heap.Push(pq,  &entry{i: i, j: j+1, priority: nums1[i] + nums2[j+1]})
		}
	}
	return res
}


func print(pq *Pq){
	for _, x := range *pq {
		fmt.Printf("is %v\n", x)
	}
	fmt.Println()
}


func main(){
	nums1 := []int{1,1,2}
	nums2 := []int{1,2,3}
	// var i = 4
	// fmt.Println(i)
	result := kSmallestPairs(nums1, nums2, 2)
	fmt.Printf("result is %v\n", result)
	// print(result)
}

