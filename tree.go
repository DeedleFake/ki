package ki

import (
	"bufio"
	"io"
	"path"
	"slices"
	"strings"
)

type Tree struct {
	c     lazyMap[string, *Tree]
	order []string
	name  string
}

func Parse(r io.Reader) (*Tree, error) {
	var root Tree

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := path.Clean(strings.TrimSpace(s.Text()))
		root.Add(line)
	}
	return &root, s.Err()
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

func (t *Tree) Children() []*Tree {
	if len(t.c) == 0 {
		return nil
	}

	c := t.c.M()
	s := make([]*Tree, 0, len(c))
	for _, name := range t.order {
		s = append(s, c[name])
	}
	return s
}

type lazyMap[K comparable, V any] map[K]V

func (m *lazyMap[K, V]) M() map[K]V {
	if *m == nil {
		*m = make(map[K]V)
	}
	return *m
}
