# Day 7

Poker! Sort the hands by "strength".

## Part 1

This one was straight forward, leaning on the functionality provided by
`slices.Sort()`. I didn't feel like implementing my own quicksort, prefering
built-in Go features, since this is about me learning Go. The only parts that
required a little thought were:

1. Identify a hand type: a `map[rune]int` counting the instances of each card
   type, then deducing with nested `if`s
1. Comparing hands of same type: i converted the cards to hexadecimal for ease
   of comparing them.

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

My first attempt got the wrong answer. I suspect it's because when hands have
wildcards, I am overwriting the card names with the strongest card names,
messing up the later sequencing that should be using the original names with the
'J's.

Yup. That was it.

A couple fun points with this one:

### Closure

I needed two versions of the my sorting function, for parts 1 and 2. However,
the function could only accept two argument, the two items being compared. I
remembered at Hack Reactor we learned about closure to give functions the
ability to persist values outside of themselves, and thankfully Go supports
this! So I wrapped my sorting function with another function that accepted a
`usePart1Rules` bool input, then returned a sorting function that would use the
respective rule set.
### Cache

When I first ran part 2, it took a good 10 seconds to complete. Looking over the
code, I realized that it was calculating all the possible iterations every time
it compared two cards, resulting in many redundant calculations. So I expanded
the closure to include a `map[string]string` cache and it dramatically improved
performance to under 1.7 seconds for both parts.