# Day 9

Extrapolation!

## Part 1

The process seems straight forward, just gotta implement it in code. Seems like a good time for recursion, not knowing how many "layers" each sequence will have.

Yup, got the answer after debugging a few little mistakes. Highlights:

* Writing my first generic, `common.AssertTrue()`
* Expanding the `common.ExtractInts()` function to correctly parse negative numbers
* Keeping the functions common to all solutions in `solutions.go`, like `Solve()`, `parseInput()`, etc; and moving everything else to a new file called `helpers.go` to make copy-pasting new days easier
* Using `Seq.iter` again, my new best friend

## Part 2

Wow, I was bracing for something strenuous but my solution took very little adaption for Part 2.