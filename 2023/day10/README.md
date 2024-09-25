# Day 10: Pipe Maze

Given a starting point in a looping path, find the distance farthest from.

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

![alt text](image.png)

I have no idea where I got that first input. But once I got the right input, I got the right answer.

## Part 2

The problem here is to find the number of tiles enclosed within the loop. I have no idea how to do this one.

I started thinking about it; my first idea is to crawl outward without crossing a pipe:

1. pick a coordinate that I know is outside the loop
2. if the tile moved to is a pipe piece, store the direction you came from as the way you cannot continue
3. travel along the pipe, if a corner is found, rotate the forbidden direction accordingly. Starting on the `*`:

```
1      2      3      4
..*..  .....  .....  .....
.S-7.  .S*7.  .*-7.  .S-7.
.|.|.  .|.|.  .|.|.  .*.|.
.L-J.  .L-J.  .L-J.  .L-J.
.....  .....  .....  .....
```

**Fig 2.** Last move was S, found a pipe perpendicular to direction of travel, can't continue S, move E and W instead

**Fig 3.** Found a NW corner, forbidden direction rotates from S to E

It's a loop, so given the starting tile (the +):

![alt text](image-1.png)

It is effectively, in this case, a horizontal pipe, and therefore one side must be "in" and the other "out". The "outside" is the one with a path  to the edge of the field.

Givens:

* The 'S' tile is on the loop
* We can mark tiles that are in the loop to distinguish them from non-loop pipe
* We can determine which side of the loop is "in" and "out"

---

New tack. I took the advice to check the AOC subreddit for advice. Some were talking about using Pick's theorem, but it would appear it requires knowing the number of points within the polygon, the info i'm looking for. So i looked for other polygon area formulas and found the Shoelace formula:

$$A = \frac{1}{2}\sum_{i=1}^n (x_iy_{i+1}-x_{i+1}y_i) $$

Testing it on the square loop:

```
.....
.S-7.
.|.|.
.L-J.
.....
```
|     | $x$ | $y$ |
| --- | --- | --- |
| 1   | 2   | 4   |
| 2   | 2   | 2   |
| 3   | 4   | 2   |
| 4   | 4   | 4   |

$$ A = \frac{1}{2}((2*2-2*4) + (2*2-2*4) + (4*4-4*2) + (4*4-4*2)) $$
****
$$ A = \frac{1}{2}(4-8 + 4-8 + 16-8 + 16-8) $$

$$ A = \frac{1}{2}(-4 - 4 + 8 + 8) $$

$$ A = \frac{1}{2}(16 - 8) $$

$$ A = 4 $$

It is 2x2, so 4 looks good. Then, to get number of inner integer points, I needs omething like this:

| Dimensions     | $x$ | $f(x)$  |
| -------------- | --- | ------- |
| 2x2            | 4   | 1       |
| 2x3            | 6   | 2       |
| 1x12, 2x6, 3x4 | 12  | 0, 4, 6 |

Maybe I can't find the number of inner points just from the area, though people are saying it can be done. What if we used Shoelace to get the area, then used Pick's theorem to get inner points?

$$A=i+\frac{b}{2}-1$$

Where $i$ is interior points, $b$ is integer points on the boundary. Plugging in the values from the simple box example gives us

$$4=i+\frac{8}{2}-1$$

$$5-4=i=1$$

And that is correct!

---

I just tried implementing it, and it got the right answer. Phew, that took me some time.
