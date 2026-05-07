## N-queens problem

Placing $n$ chess queens on an $n \times n$ chessboard so that no two queens threaten each other.

### Different ways to solve

- #### Brute force over all board configurations

    + Time complexity: $2^{n^2}$
    + State space: $n^{n^2}$
- #### Compact state representation (array of columns)
    + **Idea**: checking a row while there is already a queen is placed in that row is unnecessary
    + State space: $n^n$
    + We can make the state space much smaller by **pruning**

