## Backtracking

### When to use?

- when we want all the solutions.
- when we face multiple decisions, and we don't have enough information about the future.
- each decision, leads to new decisions.

### Examples

- Maze solving
- 4 color problem

### How it works?

- a depth-first traversal of the state space tree
- when we hit a dead end we come back and go down into another path
- the entire state space will be searched
- time complexity is exponential
- one way to make it more efficient is pruning

