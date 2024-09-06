# Advent of Code 2023
## Day 1

To run the solution, change into this directory and run:

```
❯ go run .
Solution Part 1: 53974 File: ./input.txt
Solution Part 2: 52840 File: ./input.txt
```

To run tests:

```
❯ go test
PASS
ok      aoc2023/01      0.258s
```

Part 2 was very tricky because Golang does not support overlapping regex matches (positive lookahead, `(?=(foo))`), so `"oneight"` was incorrectly turned into `11` instead of `18`. I had to make my own nasty function that examines the line starting with the last character, then the last two, etc, until it can parse a number out.