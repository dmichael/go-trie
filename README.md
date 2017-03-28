# Trie
A simple trie in Golang

This implementation is derived (not transcribed) from the Java in the excellent Algorithms by Robert Sedgewick and Kevin Wayne.

## Usage

If you would like to use this library

```go
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

```


```
$ go run main.go
# => [she shells]
```
## Constraints

* Radix is hardcoded to 256 (ASCII)
* Case-sensitive
