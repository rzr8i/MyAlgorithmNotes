## Company Party Problem (Max Independent Set on a Tree)

### Problem
- Company hierarchy is a **tree**: nodes = employees, edges = direct boss–subordinate.
- Constraint: No employee attends with their **direct upper hand** (boss).  
  Equivalently: In the tree, **no two adjacent nodes** can both attend.
- Goal: Maximize total number of attendees.

---

### DP on Trees

For each node `u`, define two states:

- `dp[u][0]` = max attendees in `u`'s subtree when **`u` does NOT attend**.
- `dp[u][1]` = max attendees in `u`'s subtree when **`u` attends**.

---

### Recurrence

Let `children(u)` be the set of `u`'s direct subordinates.

- **If `u` attends** → none of its children can attend (otherwise they'd be with their boss).  
  So we add 1 for `u` plus the best of each child's *not‑attending* state:


  $$dp[u][1] = 1 + \sum_{v \in children(u)} dp[v][0]$$

- **If `u` does NOT attend** → each child is free to attend or not, independently.  
  We take the maximum of the two options for each child:

  $$dp[u][0] = \sum_{v \in children(u)} \max(dp[v][0],\ dp[v][1])$$

---

### Base case (leaf node, no children)
$$
dp[leaf][1] = 1 \qquad dp[leaf][0] = 0
$$

---

### Answer
Let `root` be the CEO (root of the tree).  
Maximum attendees =
$$\max(dp[root][0],\ dp[root][1])$$

---

### Complexity
- **Time**: O(N) – visit each node once.
- **Space**: O(N) for storing DP values + recursion stack.

---

### Example

Tree:  
```
    A
   / \
  B   C
  |
  D
```

Compute bottom‑up:

- Leaves C, D:  
  `dp[C][1]=1, dp[C][0]=0`  
  `dp[D][1]=1, dp[D][0]=0`

- Node B:  
  `dp[B][1] = 1 + dp[D][0] = 1`  
  `dp[B][0] = max(dp[D][0], dp[D][1]) = 1`

- Node A:  
  `dp[A][1] = 1 + dp[B][0] + dp[C][0] = 1+1+0 = 2`  
  `dp[A][0] = max(dp[B][0], dp[B][1]) + max(dp[C][0], dp[C][1]) = 1+1 = 2`

Answer: `max(2,2) = 2` (e.g., invite {A, D} or {B, C}).

---

### Reconstruction

To recover which employees are invited, store a `choice` for each `dp[u][state]` and backtrack.

