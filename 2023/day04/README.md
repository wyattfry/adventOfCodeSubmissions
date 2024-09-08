# Day 4

## Part 2

After sketching on paper for a few hours, i arrived at the following algo. Using the example input would look like this:

|Winners|Card|Count|Step 1|Step 2|Step 3|Step 4|
|--|--|--|--|--|--|--|
|4|a|1|    | |
|2|b|1|+1=2| |
|2|c|1|+1=2|+2=4|
|1|d|1|+1=2|+2=4|+4=8|
|0|e|1|+1=2| |+4=6|+8=14|
|0|f|1|    | |

1. Card __a__ has four winners, so distance is four (b,c,d,e), and amount to add is the count of __a__, which is 1
1. Card __b__ has two winners, its count is 2, so add 2 to c, d
1. Card __c__ has two winners, its count is 4, add 4 to d, e
1. Card __d__ has one winner, its count is 8, add 8 to e
1. Cards __e__ and __f__ have zero winners