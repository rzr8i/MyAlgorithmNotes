## Subset with sum of M

- If we only have positive numbers, we can prune when `weight >= M`
- We can sort the numbers in preprocessing step to be able to do more pruning. (Because it has exponential time complexity, adding a polynomial preprocessing step to be able to do more pruning is actually useful)
- Also we can do more pruning by checking if `weight+total >= M`. because if it returns `false`, it means even if we include all the remaining numbers, there is no way to reach `M`, so we better prune that path and go to another one.

