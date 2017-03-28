# Trie
A simple trie in Golang

This implementation is derived (not transcribed) from the Java version as detailed in the excellent [Algorithms](http://algs4.cs.princeton.edu/home/) by Robert Sedgewick and Kevin Wayne. I recommend this text without reservation. It is detailed, thorough, and very affordable.

## Usage

If you would like to use this package, you may do something as follows. Please be aware this package may not be production ready and the implementation is incomplete.

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

## TODO

- [ ] Complete the GoDocs
- [ ] Add pattern matching funcs


