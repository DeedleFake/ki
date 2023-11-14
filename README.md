ki
==

ki is a small library and command to convert a list of paths into a tree.

Example
-------

```bash
$ tar -c . | tar -t | ki
a
  b
    three.txt
  two.txt
one.txt
```
