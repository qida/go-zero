package stringx

import (
	"sync"
)

type node struct {
	m        *sync.RWMutex
	children map[rune]*node
	end      bool
}

func (n *node) add(word string) {
	n.m.Lock()
	defer n.m.Unlock()
	chars := []rune(word)
	if len(chars) == 0 {
		return
	}
	nd := n
	for _, char := range chars {
		if nd.children == nil {
			child := new(node)
			nd.children = map[rune]*node{
				char: child,
			}
			nd = child
		} else if child, ok := nd.children[char]; ok {
			nd = child
		} else {
			child := new(node)
			nd.children[char] = child
			nd = child
		}
	}
	nd.end = true
}
