# Advent of Code 2023

![Unit Tests](https://github.com/wyattfry/adventOfCodeSubmissions/actions/workflows/go-test.yaml/badge.svg)

These puzzle solutions are in Golang. They can be run by supplying day's number. For example, to run the solution for day 1:

```
$ cd adventOfCodeSubmissions/2023
$ go run . 1
Solution Part 1: 53974 (File: ./day01/input.txt)
Solution Part 2: 52840 (File: ./day01/input.txt)
```

There are unit tests that run in a GitHub Action, or can be manually run like this:

```
$ cd adventOfCodeSubmissions/2023
$ go test ./...
?       aoc     [no test files]
ok      aoc/common      (cached)
ok      aoc/day01       (cached)
ok      aoc/day02       (cached)
ok      aoc/day03       (cached)
ok      aoc/day04       (cached)
ok      aoc/day05       0.238s
```
