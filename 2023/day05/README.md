# Day 5

## Part 1

On first reading the problem I quickly got lost. But reading it a couple more times, there's actually little to it. On paper I arrived at the algo:

```
if source >= sourceStart and source <= sourceStart + range
    source += destinationStart - sourceStart
```
Example,
```
source = 79
ds = 52
ss = 50
r = 49
if (true)
  source = 79 + (52 - 50) = 81
```

Except there was a bug that didn't emerge until...

## Part 2

You may have already spotted it:

```
if source >= sourceStart and source <= sourceStart + range
                                     ^
```

It should have been `<`.

Part 2 was the first time this year where running the solution took a noticeable amount of time, about two minutes on my M3 macbook. So I added a progress bar, this module's first non-standard-library package!