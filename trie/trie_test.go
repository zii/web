package trie

import (
	"fmt"
	"testing"
)

func Test1(_ *testing.T) {
	t := new(RuneTrie)
	for _, s := range []string{"a", "abc"} {
		t.Insert([]rune(s))
	}
	fmt.Println("match:", t.LongestPrefix([]rune("abc")))
}
