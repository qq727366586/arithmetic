package main

import "fmt"

// zh: 前缀树
// en: Trie

type Trie struct {
	pass int
	end  int
	next []*Trie
}

func NewTrie() *Trie {
	return &Trie{
		pass: 0,
		end:  0,
		next: make([]*Trie, 26, 26),
	}
}

func (t *Trie) Insert(str string) {
	root := t
	root.pass++
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		if root.next[index] == nil {
			root.next[index] = NewTrie()
		}
		root = root.next[index]
		root.pass++
	}
	root.end++
}

// str 出现几次
func (t *Trie) Search(str string) int {
	if str == "" {
		return 0
	}
	root := t
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		if root.next[index] == nil {
			return 0
		}
		root = root.next[index]
	}
	return root.end
}

// prefix 出现几次
func (t *Trie) SearchPrefix(str string) int {
	if str == "" {
		return 0
	}
	root := t
	bt := []byte(str)
	for _, v := range bt {
		index := v - 'a'
		if root.next[index] == nil {
			return 0
		}
		root = root.next[index]
	}
	return root.pass
}

func (t *Trie) Delete(str string) {
	if t.Search(str) != 0 {
		root := t
		root.pass--
		bt := []byte(str)
		for _, v := range bt {
			index := v - 'a'
			root.next[index].pass--
			if root.next[index].pass == 0 {
				root.next[index] = nil
				return
			}
			root = root.next[index]
		}
		root.end--
	}
}

func main() {
	node := NewTrie()
	node.Insert("abc")
	node.Insert("cbc")
	node.Insert("nbc")
	node.Insert("nbco")

	n := node.Search("cbc")
	fmt.Println(n)

	n = node.SearchPrefix("nb")
	fmt.Println(n)

	node.Delete("nbc")
	n = node.SearchPrefix("nb")
	fmt.Println(n)
}
