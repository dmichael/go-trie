package trie

import (
	"reflect"
	"testing"
)

var words = []string{"coffee", "cough"}

var testcases = []struct {
	input    []string // input
	expected []string // expected result
}{
	{words, words},
}

func CreateTrie(words []string) *Trie {
	dict := NewTrie()
	for _, word := range words {
		dict.Put(word)
	}
	return dict
}

func TestKeys(t *testing.T) {
	for _, tt := range testcases {
		dict := NewTrie()
		for _, word := range tt.input {
			dict.Put(word)
		}
		actual := dict.Keys()

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Keys(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}

func TestPrefix(t *testing.T) {
	trie := CreateTrie([]string{
		"hack",
		"hackerrank",
		"hawkalougie",
	})

	actual := trie.KeysWithPrefix("hac")
	expected := []string{"hack", "hackerrank"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("KeysWithPrefix(): expected %v, actual %v", actual, expected)
	}

}
