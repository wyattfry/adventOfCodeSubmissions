# Day 7

Poker! Sort the hands by "strength".

## Part 1

This one was straight forward, leaning on the functionality provided by `slices.Sort()`. I didn't feel like implementing my own quicksort, prefering built-in Go features, since this is about me learning Go. The only parts that required a little thought were:

1. Identify a hand type: a `map[rune]int` counting the instances of each card type, then deducing with nested `if`s
1. Comparing hands of same type: i converted the cards to hexadecimal for ease of comparing them.

| String | Decimal | Hex |
| ------ | ------- | --- |
| 2-9    | 2-9     | 2-9 |
| T      | 10      | a   |
| J      | 11      | b   |
| Q      | 12      | c   |
| K      | 13      | d   |
| A      | 14      | e   |

## Part 2

The twist is the `J` cards are not jacks but jokers, i.e. wildcards.

