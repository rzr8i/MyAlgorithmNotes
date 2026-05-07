## Dynamic Programming

### When to use?

- when a problem can be broken down into *overlapping subproblems* (same subproblem appears many times)
- when the problem has optimal substructure (optimal solution can be built from optimal solutions of subproblems)

### Examples

- Fibonacci sequence
- Knapsack problem
- Longest common subsequence

### How it works?

- solve subproblems once and store their results.
- avoid recomputing the same subproblem over and over.
- time complexity is usually polynomial (much better than exponential brute force)
- space complexity can be optimized (e.g., keep only last two rows instead of full table).
