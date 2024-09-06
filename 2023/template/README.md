# Advent of Code
## Golang Template

This is a basic way to organize your code for writing solutions for Advent of
Code. Each day has its own folder that contains the following:

1. *.go files, one of which has a function called `Solve()` (with captial "S" so it's accessible from other
folders/packages)
1. an `input.txt` file with the puzzle input
1. the `readlines()` function must point to the right filepath for that day:
```
package day02
...         ^
lines := readlines("./day02/input.txt")
                          ^
```

Each day's `Solve()` func can be called from the root level
`main.go`. To run, change to the root directory and run `go run .`.