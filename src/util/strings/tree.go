package strings

import "sync"

type StringBTreeNode[T comparable] struct {
	hasValue bool
	children map[T]*StringBTreeNode[T]
}

type StringBTree[T comparable] struct {
	mu   sync.RWMutex
	root *StringBTreeNode[T]
}

func NewStringBTree[T comparable]() *StringBTree[T] {
	return &StringBTree[T]{root: nil}
}

func (n *StringBTreeNode[T]) insert(s []T) {
	if len(s) == 0 {
		n.hasValue = true
		return
	}

	r := s[0]
	child, ok := n.children[r]
	if !ok {
		child = &StringBTreeNode[T]{
			children: make(map[T]*StringBTreeNode[T]),
		}
		n.children[r] = child
	}

	child.insert(s[1:])
}

func (t *StringBTree[T]) Insert(s []T) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.root == nil {
		t.root = &StringBTreeNode[T]{
			children: make(map[T]*StringBTreeNode[T]),
		}
	}

	t.root.insert(s)
}

func (n *StringBTreeNode[T]) search(s []T) bool {
	if n == nil {
		return false
	}

	if len(s) == 0 {
		return n.hasValue
	}

	r := s[0]
	child, ok := n.children[r]
	if !ok {
		return false
	}

	return child.search(s[1:])
}

func (t *StringBTree[T]) Search(s []T) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.root == nil {
		return false
	}

	return t.root.search(s)
}

func (n *StringBTreeNode[T]) delete(s []T) bool {
	if n == nil {
		return false
	}

	if len(s) == 0 {
		if n.hasValue {
			n.hasValue = false
			return true
		}
		return false
	}

	r := s[0]
	child, ok := n.children[r]
	if !ok {
		return false
	}

	if child.delete(s[1:]) {
		if len(child.children) == 0 && !child.hasValue {
			delete(n.children, r)
		}
		return true
	}

	return false
}

func (t *StringBTree[T]) Delete(s []T) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.root == nil {
		return
	}

	t.root.delete(s)
}

func (n *StringBTreeNode[T]) walk(s []T, f func([]T)) {
	if n == nil {
		return
	}

	if n.hasValue {
		f(s)
	}

	for r, child := range n.children {
		child.walk(append(s, r), f)
	}
}

func (t *StringBTree[T]) Walk(f func([]T)) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.root == nil {
		return
	}

	t.root.walk([]T{}, f)
}
