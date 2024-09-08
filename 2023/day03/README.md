# Day 3
## Part 2

Struggling on how to solve, how to break the problem into chunks.

What would be easy?

I can easily make indicies of part numbers and asterisks:

```json

.2.
.*.
.4.

.12
.*.
.7.

numbers = {
    1: 2,
    8: 4
}

asterisks = [
    5
]
```

But then what?

Or, how would i get all adjacent numbers at any index?
1. check all around. If no numbers, done.
2. if digit, expand out until non digits are found

Or,
1. getSurroundingChars(index=4, [N,NE,E,SE,S,SW,W,NW]) // location of '*'
   > 2...4...

   getSurroundingChars(index=1, [E,W]) // location of '2'
   --..*..- // no digits at E or W

   getSC(i=7, [E,W]) // location of '4'
   *..---..

I started making a recursive function but gave up. Then i got an idea: build some data structures that were more conducive to solving the problem. Namely:

1. find (row, col) of asterisks
2. calculate the (row, col)s surrounding one of those asterisk (row, col)
3. make a map where the keys are "row,col" of single digits and the values are an ID for the number
4. make a map where the keys are the number IDs, and the values are the number

With those four things, the problem is easily solved. And, you only have to iterate over the schematic once.