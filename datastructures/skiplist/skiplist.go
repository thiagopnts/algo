package skiplist

import (
	"math/rand"
)

type SkipList struct {
	head   *node
	levels int
}

type node struct {
	next  []*node
	value int
}

func newNode(value, level int) *node {
	return &node{value: value, next: make([]*node, level)}
}

func New() *SkipList {
	return &SkipList{head: newNode(0, 33), levels: 1}
}

func (s *SkipList) Insert(value int) {
	var level int = 0
	for i := rand.Int(); (i & 1) == 1; i >>= 1 {
		level++
		if level == s.levels {
			s.levels++
			break
		}
	}
	n := newNode(value, level+1)
	current := s.head
	for i := s.levels - 1; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			if current.next[i].value > value {
				break
			}
		}
		if i <= level {
			n.next[i] = current.next[i]
			current.next[i] = n
		}
	}
}

func (s *SkipList) Contains(value int) bool {
	current := s.head
	for i := s.levels - 1; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			if current.next[i].value > value {
				break
			}
			if current.next[i].value == value {
				return true
			}
		}
	}
	return false
}

func (s *SkipList) Remove(value int) bool {
  current := s.head
  found := false
  for i := s.levels - 1; i >= 0; i-- {
    for ; current.next[i] != nil; current = current.next[i] {
        if current.next[i].value == value {
            found = true
            current.next[i] = current.next[i].next[i]
            break
        }
        if current.next[i].value > value {
            break
        }
    }
  }
  return found
}

