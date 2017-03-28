package trie

import "fmt"

const radix = 256 // ASCII [a-z,A-Z]

type Trie struct {
	root  *Node
	words int
}

type Node struct {
	Value int
	links []*Node
}

func NewTrie() *Trie {
	return &Trie{newNode(), 0}
}

// NewNode initializes an empty Node with the appropriate null links to child nodes
func newNode() *Node {
	return &Node{
		links: make([]*Node, radix),
	}
}

func (trie *Trie) Get(word string) *Node {
	return get(trie.root, word, 0)
}

func get(node *Node, word string, d int) *Node {
	if node == nil {
		return nil
	}

	if d == len(word) {
		return node
	}

	letter := word[d]

	return get(node.links[letter], word, d+1)
}

// Put adds a word to the Trie
func (trie *Trie) Put(word string) *Node {
	trie.words += 1
	return put(trie.root, word, trie.words, 0)
}

// put is a recursive function that traverses the Trie's nodes creating links as necessary
func put(node *Node, word string, value int, d int) *Node {
	// Recursive calls traverse nodes which may be null from initialization
	if node == nil {
		node = newNode()
	}

	// If the cursor is the length of the word inserted, store a marker
	// indicating that this location is the end of the word
	if d == len(word) {
		node.Value = value
		return node
	}

	// The ASCII byte code is used as the index to the array
	// This has it functioning a little like a map (which may be more efficient)
	letter := word[d]

	node.links[letter] = put(node.links[letter], word, value, d+1)
	return node
}

func (trie *Trie) Keys() []string {
	return trie.KeysWithPrefix("")
}

func (trie *Trie) KeysWithPrefix(prefix string) []string {
	q := make([]string, 0)
	return collect(get(trie.root, prefix, 0), prefix, q)
}

// collect is a recursive function that collects words in the trie
func collect(node *Node, prefix string, q []string) []string {
	if node == nil {
		return q
	}

	if node.Value > 0 {
		q = append(q, prefix)
	}

	for c := 0; c < radix; c++ {
		q = collect(node.links[c], fmt.Sprintf("%s%c", prefix, c), q)
	}
	return q
}
