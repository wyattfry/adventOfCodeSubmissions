# Day 10: Pipe Maze

Given a starting point in a circular path, find the distance farthest from.

## Part 1

Sketching out an algo. Essentially, measure how long the loop is by stepping through it until we return to the starting "tile." The first problem I've run into was knowing the direction of travel, as each pipe has exactly two directions. My approach I'll first try is to "burn bridges," that is, once leaving a tile, replace it with ground: `.`

One of the examples:

```
.....
.S-7.
.|.|.
.L-J.
.....
```

But let's assume we are starting mid-path because bootstrapping the stepping may require some different handling.

```
.....
..-7.
.|.|.
.L-J.
.....
```

So let's assume the starting tile `S` has been replaced with ground, and we are now on the second row, third tile, `-`.

1. Check the kind of pipe on current tile: it's horizontal `-` so we know the next tile must be either to the east or west.
1. Check first potential next tile to the west, it must be a pipe with a connection on its east side. In this case, it's not, try next
1. Check tile to east if it has a west connector: it does. If no connecting tile is found, then we have returned to the beginning tile.
1. Next tile is found, store current location in a temporary var, change current location to new tile, overwrite old tile with ground
1. Return to step 1

...it worked! At first I didn't get the right answer. I checked my `input.txt` and noticed there were two `S` tiles, while the description implied that there would only be one. I made an HTML visualization of the loop that is fun to watch, and both loops were fairly small. I tried re-copying the input from the AoC website, it was very different. It had only one `S`, and the loop was this giant doughnut that took up most of the grid. 

![Giant Doughnut Visualization](image.png)

I have no idea where I got that first input. But once I got the right input, I got the right answer.

## Part 2

The problem here is to find the number of tiles enclosed within the loop. I have no idea how to do this one.