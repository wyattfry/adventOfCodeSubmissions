# Day 11: Expanding Galaxies

Find the shortest distances between each galaxy in expanding space

## Part 1

I wonder if rather than operating on a string array of the input lines, it might be better to make a slice of coordinates:

```
#..
...
..#
```

Would become

```
[
    {0, 0},
    {2, 2},
]
```

What would expanding it look like? Arrows indicate "doubled" or added rows or columns.

```
    v
  #...
  ....
> .... <
  ...#
    ^
```

And in the slice:

```
[
    {col: 0, row: 0},
    {col: 3, row: 3},
]
```

But by what algorithm?

1. scan for empty column
2. when one is found, increment by one all column values that are greater than the empty column
3. repeat for rows

In this example:

1. Is column 0 empty? Iterate through slice, first galaxy is in col 0, skip to next
2. Is column 1 empty? Yes, no location has col: 1
3. Iterate over slice, if any galaxy has a col > 1, increment it by one
4. second galaxy has col: 2, increment to 3
5. repeat for rows

It might be even better to represent the data as a 2D int slice, where the outer index is the col, and the inner slice is the rows

```
[
    [0],    # col 0
    [],     # col 1
    [2],    # col 2
]
```

Maybe it's not worth the trouble.

---

I was thinking about storing the data in one shape, then accessing it in another.

I was thinking a map[row][]col... and while that's great for finding which rows are not occupied, it's a pain to find columns.

Here's the next idea, use two structures. Modify the map to this:

```
map[int]map[int]int{
  row_0: {
    col_0: id_0
  },
  row_3: {
    col_2: id_1
  },
  row_4: {
    col_5: id_2
  }
}
```

Where these rows and columns are the original locations, then have a second structure:

```
map[int]location{
  id_0: {row_0, col_0},
  id_1: {row_3, col_2},
  id_1: {row_4, col_5},
}
```

...but I'd need two `map[int]map[int]int`s, one with columns first, the other with rows first. That's a lot of memory.

What if we declare an int variable for each row and column to store the number of galaxies in each. Then we iterate over the input, incrementing the count as we go, and keep that `id: location` map. That might work.

Hmm. Not quite. Ok, instead of just counts, it's a slice of galaxy IDs.

## Part 2

