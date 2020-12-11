
package main 

import "fmt"

// @lc code=start
type Trie struct {
	isWord   bool
	children [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this
	for index, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			cur.children[n] = new(Trie)
		}
		
		cur = cur.children[n]

		if index == len(word) - 1 {
			cur.isWord = true
		} 
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for index, c := range word {
		
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		
		if index == len(word) -1 {
			return true
		}

		cur = cur.children[n]
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	position := 0
	for index, c := range prefix {
		
		n := c - 'a'
		if cur.children[n] == nil {
			break
		}
		
		position++
		if index == len(prefix) -1 {
			break		
		}

		cur = cur.children[n]
	}
	return position == len(prefix) && cur.isWord == true
}


func main(){
	obj := Constructor();
	word1 := "hello"
	obj.Insert(word1);
	param_11 := obj.Search(word1);
	Parma_12 := obj.StartsWith(word1)
	fmt.Printf("search is %t, startsWith is %t\n", param_11, Parma_12)

	word2 := "hex"
	param_21 := obj.Search(word2)
	param_22 := obj.StartsWith(word2)
	fmt.Printf("search is %t, startsWith is %t\n", param_21, param_22)
	// param_3 := obj.StartsWith(prefix);

	word3 := "he"
	param_31 := obj.Search(word3)
	param_32 := obj.StartsWith(word3)
	fmt.Printf("search is %t, startsWith is %t\n", param_31, param_32)

	obj.Insert(word2)
	param_31 = obj.Search(word2)
	param_32 = obj.StartsWith(word2)
	fmt.Printf("search is %t, startsWith is %t\n", param_31, param_32)

}


