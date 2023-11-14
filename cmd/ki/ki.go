package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"deedles.dev/ki"
)

func printTree(t *ki.Tree, depth int) {
	pre := strings.Repeat("  ", depth)
	fmt.Printf("%v%v\n", pre, t.Name())
	t.Children(func(c *ki.Tree) bool {
		printTree(c, depth+1)
		return true
	})
}

func main() {
	flag.Usage = func() {
		arg0 := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %v <paths...>\n", arg0)
		fmt.Fprintf(os.Stderr, "       command-producing-paths | %v [paths...]\n", arg0)
		fmt.Fprintf(os.Stderr, "       %v [paths...] < file-with-list-of-paths\n", arg0)
	}
	flag.Parse()

	tree, err := ki.Parse(os.Stdin)
	if err != nil {
		slog.Error("parse tree", "err", err)
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		tree.Add(arg)
	}
	if tree.NumChildren() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	printTree(tree, 0)
}
