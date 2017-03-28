package main

import (
	"fmt"

	"github.com/dmichael/go-trie/trie"
)

var dict = []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

func main() {
	t := trie.NewTrie()
	for _, word := range dict {
		t.Put(word)
	}

	matches := t.KeysWithPrefix("she")
	fmt.Println(matches)
}
