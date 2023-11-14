// ki is a tool for converting lists of paths into a tree.
//
// Usage:
//
//	ki <paths...>
//	tar -tf file.tar | ki
//	ki < file.txt
//
// ki takes a list of slash-separated paths and produces a nicely
// formatted, alphabetically sorted tree structure of the elements of
// the path. It can read its input list either from standard input,
// one per line, or as command-line arguments, or both.
package main
