// Package ki creates tree structures from lists of file paths.
package ki

import (
	"bufio"
	"io"
	"path"
	"slices"
	"strings"
)

// A Tree is a tree of paths. Each Tree instance is a directory in
// that path tree.
type Tree struct {
	c     lazyMap[string, *Tree]
	order []string
	name  string
}

// Parse parses a nweline-separated list of paths and builds a tree
// from it.
func Parse(r io.Reader) (*Tree, error) {
	var root Tree
	err := root.Parse(r)
	return &root, err
}

// Parse parses a newline-separated list of paths and adds it to the
// tree rooted at t.
func (t *Tree) Parse(r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := path.Clean(strings.TrimSpace(s.Text()))
		t.Add(line)
	}
	return s.Err()
}

func (t *Tree) addPath(parts []string) {
	if len(parts) == 0 {
		return
	}

	t.child(parts[0]).addPath(parts[1:])
}

func (t *Tree) child(name string) *Tree {
	c := t.c.M()
	v, ok := c[name]
	if ok {
		return v
	}

	v = &Tree{name: name}
	c[name] = v
	t.insert(name)
	return v
}

func (t *Tree) insert(name string) {
	i, _ := slices.BinarySearch(t.order, name)
	t.order = slices.Insert(t.order, i, name)
}

// Add adds a single path to the tree rooted at t.
func (t *Tree) Add(p string) {
	if path.IsAbs(p) {
		p = p[1:]
	}
	if (p == "") || (p == ".") {
		return
	}

	parts := strings.Split(p, "/")
	t.addPath(parts)
}

// Children yields the child trees of t in ascending alphabetical
// order.
func (t *Tree) Children(yield func(*Tree) bool) {
	if len(t.c) == 0 {
		return
	}

	c := t.c.M()
	for _, name := range t.order {
		if !yield(c[name]) {
			return
		}
	}
	return
}

// NumChildren returns the number of t's child trees.
func (t *Tree) NumChildren() int {
	return len(t.c)
}

type lazyMap[K comparable, V any] map[K]V

func (m *lazyMap[K, V]) M() map[K]V {
	if *m == nil {
		*m = make(map[K]V)
	}
	return *m
}
