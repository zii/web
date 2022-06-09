package trie

type RuneTrie struct {
	root *node
}

type node struct {
	leaf bool
	next map[rune]*node
}

func (this *RuneTrie) Insert(key []rune) {
	if this.root == nil {
		this.root = new(node)
	}
	if len(key) <= 0 {
		return
	}
	cur := this.root
	for _, c := range key {
		if cur.next == nil {
			cur.next = make(map[rune]*node)
		}
		if n, ok := cur.next[c]; ok {
			cur = n
		} else {
			cur.next[c] = new(node)
			cur = cur.next[c]
		}
	}
	cur.leaf = true
}

func (this *RuneTrie) LongestPrefix(key []rune) int {
	/* Finds the longest prefix of a key with a value. Return length. */
	n := 0
	cur := this.root
	for _, c := range key {
		if cur == nil {
			break
		}
		if d, ok := cur.next[c]; ok {
			cur = d
			n++
		} else {
			break
		}
	}
	if cur != nil && !cur.leaf {
		return 0
	}
	return n
}
